import React, { useState, useEffect } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { Dialog } from 'primereact/dialog';
import Swal from 'sweetalert2';
import 'primereact/resources/themes/saga-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import PaymentComponent from './PaymentComponent';

const DataTableComponent = () => {
  const [customers, setCustomers] = useState([]);
  const [filter, setFilter] = useState('');
  const [newCustomer, setNewCustomer] = useState({
    marketingID: '',
    date: '',
    cargoFee: '',
    totalBalance: '',
  });
  const [editingCustomer, setEditingCustomer] = useState(null);
  const [isDialogVisible, setDialogVisible] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [first, setFirst] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showPayment, setShowPayment] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      setError(null);
      try {
        const response = await fetch(`${process.env.REACT_APP_API_URL}/penjualan`);
        if (!response.ok) throw new Error('Network response was not ok');
        const data = await response.json();
        setCustomers(data);
      } catch (error) {
        setError('Failed to fetch customer data.');
        console.error('Error fetching data:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewCustomer((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const openDialog = (customer = null) => {
    setEditingCustomer(customer);
    setNewCustomer(customer || {
      marketingID: '',
      date: '',
      cargoFee: '',
      totalBalance: '',
    });
    setIsEditing(Boolean(customer));
    setDialogVisible(true);
  };

  const handleCustomerAction = async (method) => {
    const apiUrl = `${process.env.REACT_APP_API_URL}/penjualan${isEditing ? `/${editingCustomer.id}` : ''}`;
    
    // Ensure cargoFee and totalBalance are parsed to numbers
    const payload = {
      marketingID: parseInt(newCustomer.marketingID, 10), // parseInt with base 10
      date: newCustomer.date,
      cargoFee: parseFloat(newCustomer.cargoFee) || 0, // Use parseFloat and default to 0 if NaN
      totalBalance: parseFloat(newCustomer.totalBalance) || 0, // Same here
    };
  
    try {
      const response = await fetch(apiUrl, {
        method,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
  
      if (!response.ok) {
        const errorMessage = await response.text();
        throw new Error(`Failed to ${method === 'POST' ? 'add' : 'update'} customer: ${errorMessage}`);
      }
  
      const result = await response.json();
      setCustomers((prevCustomers) => isEditing
        ? prevCustomers.map((customer) => (customer.id === editingCustomer.id ? result : customer))
        : [...prevCustomers, result]
      );
  
      Swal.fire({
        title: method === 'POST' ? 'Success!' : 'Updated!',
        text: `Customer ${method === 'POST' ? 'added' : 'updated'} successfully.`,
        icon: 'success',
        confirmButtonText: 'Ok',
      });
      setDialogVisible(false);
    } catch (error) {
      Swal.fire({
        title: 'Error!',
        text: error.message,
        icon: 'error',
        confirmButtonText: 'Ok',
      });
    }
  };
  

  const handleDeleteCustomer = (customerId) => {
    Swal.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#d33',
      cancelButtonColor: '#3085d6',
      confirmButtonText: 'Yes, delete it!',
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          const response = await fetch(`${process.env.REACT_APP_API_URL}/penjualan/${customerId}`, { method: 'DELETE' });
          if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to delete customer');
          }
          setCustomers((prevCustomers) => prevCustomers.filter((customer) => customer.id !== customerId));
          Swal.fire('Deleted!', 'The customer has been deleted.', 'success');
        } catch (error) {
          Swal.fire({
            title: 'Error!',
            text: error.message || 'Failed to delete customer.',
            icon: 'error',
            confirmButtonText: 'Ok',
          });
        }
      }
    });
  };

  const filteredCustomers = customers.filter((customer) =>
    customer.TransactionNumber?.toLowerCase().includes(filter.toLowerCase())
  );

  return (
    <div>
      <h3>List Penjualan</h3>
      {error && <p className="error">{error}</p>}
      <Button label="Add New Transaction" icon="pi pi-plus" onClick={() => openDialog()} />
      <div className="p-mb-2">
        <select value={rowsPerPage} onChange={(e) => {
          setRowsPerPage(Number(e.target.value));
          setFirst(0);
        }}>
          {[5, 10, 25, 50].map((num) => (
            <option key={num} value={num}>{num}</option>
          ))}
        </select>
      </div>
      <div className="p-mb-5">
        <InputText
          value={filter}
          onChange={(e) => setFilter(e.target.value)}
          placeholder="Search by Transaction Number"
        />
      </div>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <DataTable
          value={filteredCustomers.slice(first, first + rowsPerPage)}
          showGridlines
          paginator
          rows={rowsPerPage}
          first={first}
          onPage={(e) => setFirst(e.first)}
          totalRecords={filteredCustomers.length}
          rowsPerPageOptions={[5, 10, 25, 50]}
        >
          <Column field="id" header="ID" sortable style={{ width: '5%' }} />
          <Column field="TransactionNumber" header="Transaction Number" />
          <Column field="MarketingID" header="Marketing ID" />
          <Column field="Date" header="Date" />
          <Column field="CargoFee" header="Cargo Fee" />
          <Column field="TotalBalance" header="Total Balance" />
          <Column field="GrandTotal" header="Grand Total" />
          <Column
            header="Actions"
            body={(rowData) => (
              <div>
                <Button icon="pi pi-pencil" onClick={() => openDialog(rowData)} className="p-button-warning" />
                <Button icon="pi pi-trash" onClick={() => handleDeleteCustomer(rowData.id)} className="p-button-danger" />
              </div>
            )}
          />
        </DataTable>
      )}

      {filteredCustomers.length === 0 && !loading && <p>No available options</p>}

      <Dialog
        visible={isDialogVisible}
        onHide={() => setDialogVisible(false)}
        header={isEditing ? 'Edit Customer' : 'Add New Customer'}
        footer={
          <div>
            <Button label="Cancel" icon="pi pi-times" onClick={() => setDialogVisible(false)} />
            <Button
              label={isEditing ? 'Update Customer' : 'Add Customer'}
              icon="pi pi-check"
              onClick={() => handleCustomerAction(isEditing ? 'PUT' : 'POST')}
            />
          </div>
        }
      >
        <div>
          <InputText name="marketingID" value={newCustomer.marketingID} onChange={handleInputChange} placeholder="Marketing ID" className="p-mb-2" required />
          <InputText name="date" type="date" value={newCustomer.date} onChange={handleInputChange} className="p-mb-2" required />
          <InputText name="cargoFee" value={newCustomer.cargoFee} onChange={handleInputChange} placeholder="Cargo Fee" className="p-mb-2" required />
          <InputText name="totalBalance" value={newCustomer.totalBalance} onChange={handleInputChange} placeholder="Total Balance" className="p-mb-2" required />
        </div>
      </Dialog>

      <Button label="Make a Payment" icon="pi pi-money-bill" onClick={() => setShowPayment(true)} />
      {showPayment && <PaymentComponent onPaymentSuccess={(payment) => console.log('Payment was successful:', payment)} />}
    </div>
  );
};

export default DataTableComponent;
