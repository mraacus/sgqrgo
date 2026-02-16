package sgqr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePaynowQrString(t *testing.T) {
	tests := []struct {
		name    string
		options PayNowQROptions
		prepare func()
		want    *string
		wantErr bool
	}{
		{
			name: "success - non editable",
			options: PayNowQROptions{
				MobileNumber:    "+6581010321",
				Editable:        false,
				Expiry:          "20261228",
				Amount:          "10.50",
				MerchantName:    "sgqr_test",
				ReferenceNumber: "REF123",
			},
			want:    convStringToPtr("00020101021226500009SG.PAYNOW010100211+658101032103010040820261228520400005303702540510.505802SG5909sgqr_test6009Singapore62100106REF1236304D9F1"),
			wantErr: false,
		},
		{
            name: "success - editable amount",
            options: PayNowQROptions{
                MobileNumber:    "+6581010321",
				Editable:        true,
				Expiry:          "20261228",
				Amount:          "10.50",
				MerchantName:    "sgqr_test",
				ReferenceNumber: "REF123",
            },
            want:    convStringToPtr("00020101021226500009SG.PAYNOW010100211+658101032103011040820261228520400005303702540510.505802SG5909sgqr_test6009Singapore62100106REF12363045E42"),
            wantErr: false,
        },
		{
            name: "failure - empty mobile number",
            options: PayNowQROptions{
                MobileNumber:    "",
                Editable:        false,
                Expiry:          "20261228",
                Amount:          "10.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - mobile number without country code",
            options: PayNowQROptions{
                MobileNumber:    "6581010321",
                Editable:        false,
                Expiry:          "20261228",
                Amount:          "10.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - mobile number with letters",
            options: PayNowQROptions{
                MobileNumber:    "+65ABC10321",
                Editable:        false,
                Expiry:          "20261228",
                Amount:          "10.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - empty expiry",
            options: PayNowQROptions{
                MobileNumber:    "+6581010321",
                Editable:        false,
                Expiry:          "",
                Amount:          "10.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - invalid expiry format",
            options: PayNowQROptions{
                MobileNumber:    "+6581010321",
                Editable:        false,
                Expiry:          "2026-12-28",
                Amount:          "10.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - empty amount",
            options: PayNowQROptions{
                MobileNumber:    "+6581010321",
                Editable:        false,
                Expiry:          "20261228",
                Amount:          "",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - invalid amount format",
            options: PayNowQROptions{
                MobileNumber:    "+6581010321",
                Editable:        false,
                Expiry:          "20261228",
                Amount:          "ABC.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "failure - negative amount",
            options: PayNowQROptions{
                MobileNumber:    "+6581010321",
                Editable:        false,
                Expiry:          "20261228",
                Amount:          "-10.50",
                MerchantName:    "sgqr_test",
                ReferenceNumber: "REF123",
            },
            want:    nil,
            wantErr: true,
        },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := GeneratePayNowQrString(tt.options)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
				if tt.want == nil {
					assert.Nil(t, got)
				} else {
					assert.Equal(t, *tt.want, got)
				}
			}
		})
	}
}

func TestGenerateSGQRString(t *testing.T) {

	tests := []struct {
		name    string
		options SGQROptions
		prepare func()
		want    *string
		wantErr bool
	}{
		{
			name: "success - uen",
			options: SGQROptions{
				ReceiverType:             ReceiverTypeUEN,
				MobileOrUENAccountNumber: "T11LL1111C",
				Editable:                 false,
				Expiry:                   "20261228",
				Amount:                   "10.50",
				SGQRID:                   "SGQR1234567890",
				MerchantName:             "sgqr_test",
				ReferenceNumber:          "REF123",
			},
			want:    convStringToPtr("00020101021226490009SG.PAYNOW010120210T11LL1111C030100408202612285114SGQR1234567890520400005303702540510.505802SG5909sgqr_test6009Singapore62100106REF1236304622F"),
			wantErr: false,
		},
		{
            name: "success - mobile number",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeMobile,
                MobileOrUENAccountNumber: "+6581010321",
                Editable:                 true,
                Expiry:                   "20261228",
                Amount:                   "25.00",
                SGQRID:                   "SGQR9876543210",
                MerchantName:             "test_merchant",
                ReferenceNumber:          "REF999",
            },
            want:    convStringToPtr("00020101021226500009SG.PAYNOW010100211+6581010321030110408202612285114SGQR9876543210520400005303702540525.005802SG5913test_merchant6009Singapore62100106REF9996304D1BD"),
            wantErr: false,
        },
		{
			name: "failure - invalid UEN",
			options: SGQROptions{
				ReceiverType:             ReceiverTypeUEN,
				MobileOrUENAccountNumber: "INVALID_UEN", // non alphanumeric
				Editable:                 false,
				Expiry:                   "20261228",
				Amount:                   "10.50",
				SGQRID:                   "SGQR1234567890",
				MerchantName:             "sgqr_test",
				ReferenceNumber:          "REF123",
			},
			want:    nil,
			wantErr: true,
		},
		{
            name: "failure - UEN too short",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL11",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "failure - UEN too long",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL1111CC",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - empty mobile number",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeMobile,
                MobileOrUENAccountNumber: "",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - mobile number with no country code",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeMobile,
                MobileOrUENAccountNumber: "81010321", // missing country code
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "failure - invalid mobile number with letters",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeMobile,
                MobileOrUENAccountNumber: "+65ABC10321",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - empty expiry date",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL1111C",
                Editable:                 false,
                Expiry:                   "",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - invalid expiry date format",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL1111C",
                Editable:                 false,
                Expiry:                   "2026-12-28",
                Amount:                   "10.50",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - invalid amount format",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL1111C",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "ABC",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "failure - empty amount",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL1111C",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "",
                SGQRID:                   "SGQR1234567890",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
		{
            name: "failure - invalid SGQR ID format",
            options: SGQROptions{
                ReceiverType:             ReceiverTypeUEN,
                MobileOrUENAccountNumber: "T11LL1111C",
                Editable:                 false,
                Expiry:                   "20261228",
                Amount:                   "10.50",
                SGQRID:                   "INVALID123",
                MerchantName:             "sgqr_test",
                ReferenceNumber:          "REF123",
            },
            want:    nil,
            wantErr: true,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}
			got, err := GenerateSGQRString(tt.options)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
				if tt.want == nil {
					assert.Nil(t, got)
				} else {
					assert.Equal(t, *tt.want, got)
				}
			}
		})
	}
}