package models

import "time"

// Pembayaran represents a payment record.
type Pembayaran struct {
	ID            int       `json:"id"`
	MarketingID   int       `json:"marketing_id"`
	Amount        float64   `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
}

// PembayaranAngsuran represents an installment payment record.
type PembayaranAngsuran struct {
	ID                int       `json:"id"`
	PembayaranID      int       `json:"pembayaran_id"`
	AngsuranKe        int       `json:"angsuran_ke"`
	JumlahAngsuran    float64   `json:"jumlah_angsuran"`
	TanggalPembayaran time.Time `json:"tanggal_pembayaran"`
}
