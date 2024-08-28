import React, { useEffect, useState } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import Swal from 'sweetalert2';

const KomisiPage = () => {
  const [komisiData, setKomisiData] = useState([]);
  const [marketingData, setMarketingData] = useState({});

  useEffect(() => {
    const fetchKomisiData = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_API_URL}/komisi`);
        const data = await response.json();
        setKomisiData(data);
      } catch (error) {
        Swal.fire({
          title: 'Error!',
          text: 'Failed to fetch komisi data.',
          icon: 'error',
          confirmButtonText: 'Ok',
        });
      }
    };

    const fetchMarketingData = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_API_URL}/marketing`);
        const data = await response.json();

        // Map marketing ID to name
        const marketingMap = {};
        data.forEach(marketing => {
          marketingMap[marketing.id] = marketing.name;
        });
        setMarketingData(marketingMap);
      } catch (error) {
        Swal.fire({
          title: 'Error!',
          text: 'Failed to fetch marketing data.',
          icon: 'error',
          confirmButtonText: 'Ok',
        });
      }
    };

    fetchKomisiData();
    fetchMarketingData();
  }, []);

  // Custom function to format komisi_persen as percentage
  const formatKomisiPersen = (rowData) => {
    return `${rowData.komisi_persen}%`; // Add percentage sign
  };

  // Custom function to get marketing name
  const getMarketingName = (rowData) => {
    return marketingData[rowData.marketing_id] || 'Unknown'; // Fallback if name not found
  };

  return (
    <div>
      <h3>Komisi Data</h3>
      <DataTable value={komisiData}>
        <Column field="marketing_id" header="Marketing Name" body={getMarketingName} /> {/* Use custom body function */}
        <Column field="bulan" header="Bulan" />
        <Column field="omzet" header="Omzet" />
        <Column field="komisi_persen" header="Komisi Persen" body={formatKomisiPersen} /> {/* Use custom body function */}
        <Column field="komisi_nominal" header="Komisi Nominal" />
      </DataTable>
    </div>
  );
};

export default KomisiPage;
