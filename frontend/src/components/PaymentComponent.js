import React, { useState, useEffect } from 'react';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import Swal from 'sweetalert2';

const PaymentComponent = ({ onPaymentSuccess }) => {
  const [paymentData, setPaymentData] = useState({
    marketingID: '',
    amount: '',
    paymentMethod: 'credit',
    jumlahAngsuran: 1,
  });
  const [validMarketingIDs, setValidMarketingIDs] = useState([]);
  const [inputError, setInputError] = useState('');

  useEffect(() => {
    const fetchValidMarketingIDs = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_API_URL}/marketing`);
        if (!response.ok) throw new Error('Failed to fetch marketing IDs');
        const data = await response.json();
        setValidMarketingIDs(data);
      } catch (error) {
        console.error('Error fetching valid marketing IDs:', error);
      }
    };

    fetchValidMarketingIDs();
  }, []);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setPaymentData((prev) => ({
      ...prev,
      [name]: value,
    }));
    setInputError(''); // Clear error on input change
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Check if the Marketing ID is valid
    if (!validMarketingIDs.includes(paymentData.marketingID)) {
      setInputError('The entered Marketing ID does not exist.'); // Set error message
      return; // Exit the function early
    }

    const amount = parseFloat(paymentData.amount);
    const jumlahAngsuran = parseInt(paymentData.jumlahAngsuran, 10);

    // Validate Amount and JumlahAngsuran
    if (amount <= 0 || jumlahAngsuran <= 0) {
      setInputError('Amount and Installments must be greater than zero.');
      return;
    }

    const payload = {
      ...paymentData,
      amount,
      jumlahAngsuran,
    };

    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/pembayaran?jumlah_angsuran=${jumlahAngsuran}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorMessage = await response.text();
        throw new Error(errorMessage);
      }

      const result = await response.json();
      Swal.fire({
        title: 'Success!',
        text: 'Payment processed successfully.',
        icon: 'success',
        confirmButtonText: 'Ok',
      });

      onPaymentSuccess(result);
      // Reset form state after successful payment
      setPaymentData({
        marketingID: '',
        amount: '',
        paymentMethod: 'credit',
        jumlahAngsuran: 1,
      });
    } catch (error) {
      Swal.fire({
        title: 'Error!',
        text: error.message,
        icon: 'error',
        confirmButtonText: 'Ok',
      });
    }
  };

  return (
    <div>
      <h3>Make a Payment</h3>
      <form onSubmit={handleSubmit}>
        <InputText
          name="marketingID"
          value={paymentData.marketingID}
          onChange={handleInputChange}
          placeholder="Marketing ID"
          className={`p-mb-2 ${inputError ? 'p-invalid' : ''}`} // Error class if there's an error
          required
        />
        {inputError && <small className="p-error">{inputError}</small>} {/* Show error message */}
        
        <InputText
          name="amount"
          value={paymentData.amount}
          onChange={handleInputChange}
          placeholder="Amount"
          className="p-mb-2"
          required
          type="number"
        />
        <InputText
          name="jumlahAngsuran"
          value={paymentData.jumlahAngsuran}
          onChange={handleInputChange}
          placeholder="Installments"
          className="p-mb-2"
          required
          type="number"
        />
        <div>
          <label>
            <input
              type="radio"
              name="paymentMethod"
              value="credit"
              checked={paymentData.paymentMethod === 'credit'}
              onChange={handleInputChange}
            />
            Credit
          </label>
          <label>
            <input
              type="radio"
              name="paymentMethod"
              value="cash"
              checked={paymentData.paymentMethod === 'cash'}
              onChange={handleInputChange}
            />
            Cash
          </label>
        </div>
        <Button type="submit" label="Submit Payment" icon="pi pi-check" />
      </form>
    </div>
  );
};

export default PaymentComponent;
