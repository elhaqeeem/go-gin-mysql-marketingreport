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

  useEffect(() => {
    const apiUrl = `${process.env.REACT_APP_API_URL}/penjualan`;
    const fetchData = async () => {
      setLoading(true);
      setError(null);
      try {
        const response = await fetch(apiUrl);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
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

  const openNewCustomerDialog = () => {
    setNewCustomer({
      marketingID: '',
      date: '',
      cargoFee: '',
      totalBalance: '',
    });
    setIsEditing(false);
    setDialogVisible(true);
  };

  const openEditCustomerDialog = (customer) => {
    setEditingCustomer(customer);
    setNewCustomer(customer);
    setIsEditing(true);
    setDialogVisible(true);
  };

  const handleAddCustomer = async () => {
    console.log('Attempting to add customer with data:', newCustomer);

    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/penjualan`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          marketingID: parseInt(newCustomer.marketingID), // Convert to integer
          date: newCustomer.date,
          cargoFee: parseFloat(newCustomer.cargoFee), // Ensure it's a number
          totalBalance: parseFloat(newCustomer.totalBalance), // Ensure it's a number
        }),
      });

      console.log('Response status:', response.status); // Log the status code

      if (!response.ok) {
        const errorMessage = await response.text(); // Get the error message
        console.error('Error response:', errorMessage); // Log the error response
        throw new Error(`Failed to add customer: ${errorMessage}`);
      }

      const result = await response.json(); // Parse the JSON response
      console.log('Add customer result:', result); // Log the result

      // Assuming the backend returns the new customer with its ID
      setCustomers((prevCustomers) => [...prevCustomers, result]);
      setDialogVisible(false);
      Swal.fire({
        title: 'Success!',
        text: 'Customer added successfully.',
        icon: 'success',
        confirmButtonText: 'Ok'
      });
    } catch (error) {
      console.error('Error adding customer:', error); // Log the error
      Swal.fire({
        title: 'Error!',
        text: error.message,
        icon: 'error',
        confirmButtonText: 'Ok'
      });
    }
  };

  const handleUpdateCustomer = async () => {
    try {
        // Prepare the data
        const updatedCustomerData = {
          MarketingID: parseInt(newCustomer.marketingID, 10), // Ensure it's an integer
          Date: newCustomer.date,
          CargoFee: parseFloat(newCustomer.cargoFee),
          TotalBalance: parseFloat(newCustomer.totalBalance),
          id: editingCustomer.id, // Include the ID in the body for the Go backend to use
      };
      

        // Log the payload to inspect it
        console.log("Sending data:", updatedCustomerData);

        // Make the PUT request
        const response = await fetch(`${process.env.REACT_APP_API_URL}/penjualan/${editingCustomer.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedCustomerData),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to update customer');
        }

        // Update the state with the new data
        setCustomers((prevCustomers) =>
            prevCustomers.map((customer) =>
                customer.id === editingCustomer.id ? { ...customer, ...newCustomer } : customer
            )
        );
        setDialogVisible(false);
        Swal.fire({
            title: 'Updated!',
            text: 'Customer updated successfully.',
            icon: 'success',
            confirmButtonText: 'Ok',
        });
    } catch (error) {
        Swal.fire({
            title: 'Error!',
            text: error.message || 'Failed to update customer.',
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
          const response = await fetch(`${process.env.REACT_APP_API_URL}/penjualan/${customerId}`, {
            method: 'DELETE',
          });

          if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to delete customer');
          }

          // Update the local state by filtering out the deleted customer
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

      <Button label="Add New Customer" icon="pi pi-plus" onClick={openNewCustomerDialog} />

      {/* Rows Per Page Dropdown */}
      <div className="p-mb-2">
        <select value={rowsPerPage} onChange={(e) => {
          setRowsPerPage(Number(e.target.value));
          setFirst(0); // Reset first when changing rows per page
        }}>
          <option value={5}>5</option>
          <option value={10}>10</option>
          <option value={25}>25</option>
          <option value={50}>50</option>
        </select>
      </div>

      {/* Search Filter */}
      <div className="p-mb-5">
        <InputText
          value={filter}
          onChange={(e) => setFilter(e.target.value)}
          placeholder="Search by Transaction Number"
        />
      </div>

      {/* Data Table */}
      {loading ? (
        <p>Loading...</p>
      ) : (
        <DataTable
          value={filteredCustomers.slice(first, first + rowsPerPage)}
          showGridlines
          paginator
          rows={rowsPerPage}
          first={first}
          onPage={(e) => {
            setFirst(e.first); // Update first when page changes
          }}
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
                <Button
                  icon="pi pi-pencil"
                  onClick={() => openEditCustomerDialog(rowData)}
                  className="p-button-warning"
                />
                <Button
                  icon="pi pi-trash"
                  onClick={() => handleDeleteCustomer(rowData.id)}
                  className="p-button-danger"
                />
              </div>
            )}
          />
        </DataTable>
      )}

      {filteredCustomers.length === 0 && !loading && (
        <p>No available options</p>
      )}

      {/* Customer Dialog */}
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
              onClick={isEditing ? handleUpdateCustomer : handleAddCustomer}
            />
          </div>
        }
      >
        <div>
          <InputText
            name="marketingID"
            value={newCustomer.marketingID}
            onChange={handleInputChange}
            placeholder="Marketing ID"
            className="p-mb-2"
            required
          />
          <InputText
            name="date"
            type="date"
            value={newCustomer.date}
            onChange={handleInputChange}
            className="p-mb-2"
            required
          />
          <InputText
            name="cargoFee"
            value={newCustomer.cargoFee}
            onChange={handleInputChange}
            placeholder="Cargo Fee"
            className="p-mb-2"
            required
          />
          <InputText
            name="totalBalance"
            value={newCustomer.totalBalance}
            onChange={handleInputChange}
            placeholder="Total Balance"
            className="p-mb-2"
            required
          />
        </div>
      </Dialog>
    </div>
  );
};

export default DataTableComponent;
