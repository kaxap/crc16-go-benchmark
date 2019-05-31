package crc16ccitt

import (
	"crypto/rand"
	howeyc "github.com/howeyc/crc16"
	efault "github.com/npat-efault/crc16"
	"github.com/sigurn/crc16"
	"hash"
	"testing"
)

func createRandomSlize(size int) []byte {
	s := make([]byte, size, size)
	rand.Read(s)
	return s
}

// Test against this.
var fodder100MB = createRandomSlize(100 * 1024 * 1024)
var fodder1MB = createRandomSlize(1024 * 1024)
var fodder1KB = createRandomSlize(1024)
var fodder64B = createRandomSlize(64)
var fodder24B = createRandomSlize(24)

func checksumTest(b *testing.B, h hash.Hash, fodder []byte) {
	for i := 0; i < b.N; i++ {
		h.Write(fodder)
		h.Sum(nil)
	}
}

// 100 MB benchmark
func Benchmark_CCITT_100MB(b *testing.B) {
	checksumTest(b, NewCrc16CCITT(), fodder100MB)
}
func Benchmark_HoweycIBM_100MB(b *testing.B) {
	checksumTest(b, howeyc.NewIBM(), fodder100MB)
}
func Benchmark_EfaultPPP_100MB(b *testing.B) {
	checksumTest(b, efault.New(efault.PPP), fodder100MB)
}
func Benchmark_EfaultModus_100MB(b *testing.B) {
	checksumTest(b, efault.New(efault.Modbus), fodder100MB)
}
func Benchmark_EfaultKermit_100MB(b *testing.B) {
	checksumTest(b, efault.New(efault.Kermit), fodder100MB)
}
func Benchmark_EfaultXModem_100MB(b *testing.B) {
	checksumTest(b, efault.New(efault.XModem), fodder100MB)
}
func Benchmark_JoaoJeronimo_100MB(b *testing.B) {
	checksumTest(b, newJjWrapper(), fodder100MB)
}
func Benchmark_Jiguorui_100MB(b *testing.B) {
	checksumTest(b, newJigWrapper(), fodder100MB)
}
func Benchmark_SigurnCRC16ARC_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_ARC), fodder100MB)
}
func Benchmark_SigurnCRC16_AUG_CCITT_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_AUG_CCITT), fodder100MB)
}
func Benchmark_SigurnCRC16_BUYPASS_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_BUYPASS), fodder100MB)
}
func Benchmark_SigurnCRC16_CCITT_FALSE_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CCITT_FALSE), fodder100MB)
}
func Benchmark_SigurnCRC16_CDMA2000_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CDMA2000), fodder100MB)
}
func Benchmark_SigurnCRC16_DDS_110_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DDS_110), fodder100MB)
}
func Benchmark_SigurnCRC16_DECT_R_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_R), fodder100MB)
}
func Benchmark_SigurnCRC16_DECT_X_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_X), fodder100MB)
}
func Benchmark_SigurnCRC16_DNP_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DNP), fodder100MB)
}
func Benchmark_SigurnCRC16_EN_13757_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_EN_13757), fodder100MB)
}
func Benchmark_SigurnCRC16_GENIBUS_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_GENIBUS), fodder100MB)
}
func Benchmark_SigurnCRC16_MAXIM_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MAXIM), fodder100MB)
}
func Benchmark_SigurnCRC16_MCRF4XX_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MCRF4XX), fodder100MB)
}
func Benchmark_SigurnCRC16_RIELLO_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_RIELLO), fodder100MB)
}
func Benchmark_SigurnCRC16_T10_DIF_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_T10_DIF), fodder100MB)
}
func Benchmark_SigurnCRC16_TELEDISK_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TELEDISK), fodder100MB)
}
func Benchmark_SigurnCRC16_TMS37157_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TMS37157), fodder100MB)
}
func Benchmark_SigurnCRC16_USB_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_USB), fodder100MB)
}
func Benchmark_SigurnCRC16_CRC_A_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CRC_A), fodder100MB)
}
func Benchmark_SigurnCRC16_KERMIT_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_KERMIT), fodder100MB)
}
func Benchmark_SigurnCRC16_MODBUS_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MODBUS), fodder100MB)
}
func Benchmark_SigurnCRC16_X_25_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_X_25), fodder100MB)
}
func Benchmark_SigurnCRC16_XMODEM_100MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_XMODEM), fodder100MB)
}

// 1 MB benchmark
func Benchmark_CCITT_1MB(b *testing.B) {
	checksumTest(b, NewCrc16CCITT(), fodder1MB)
}
func Benchmark_HoweycIBM_1MB(b *testing.B) {
	checksumTest(b, howeyc.NewIBM(), fodder1MB)
}
func Benchmark_EfaultPPP_1MB(b *testing.B) {
	checksumTest(b, efault.New(efault.PPP), fodder1MB)
}
func Benchmark_EfaultModus_1MB(b *testing.B) {
	checksumTest(b, efault.New(efault.Modbus), fodder1MB)
}
func Benchmark_EfaultKermit_1MB(b *testing.B) {
	checksumTest(b, efault.New(efault.Kermit), fodder1MB)
}
func Benchmark_EfaultXModem_1MB(b *testing.B) {
	checksumTest(b, efault.New(efault.XModem), fodder1MB)
}
func Benchmark_JoaoJeronimo_1MB(b *testing.B) {
	checksumTest(b, newJjWrapper(), fodder1MB)
}
func Benchmark_Jiguorui_1MB(b *testing.B) {
	checksumTest(b, newJigWrapper(), fodder1MB)
}
func Benchmark_SigurnCRC16ARC_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_ARC), fodder1MB)
}
func Benchmark_SigurnCRC16_AUG_CCITT_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_AUG_CCITT), fodder1MB)
}
func Benchmark_SigurnCRC16_BUYPASS_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_BUYPASS), fodder1MB)
}
func Benchmark_SigurnCRC16_CCITT_FALSE_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CCITT_FALSE), fodder1MB)
}
func Benchmark_SigurnCRC16_CDMA2000_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CDMA2000), fodder1MB)
}
func Benchmark_SigurnCRC16_DDS_110_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DDS_110), fodder1MB)
}
func Benchmark_SigurnCRC16_DECT_R_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_R), fodder1MB)
}
func Benchmark_SigurnCRC16_DECT_X_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_X), fodder1MB)
}
func Benchmark_SigurnCRC16_DNP_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DNP), fodder1MB)
}
func Benchmark_SigurnCRC16_EN_13757_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_EN_13757), fodder1MB)
}
func Benchmark_SigurnCRC16_GENIBUS_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_GENIBUS), fodder1MB)
}
func Benchmark_SigurnCRC16_MAXIM_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MAXIM), fodder1MB)
}
func Benchmark_SigurnCRC16_MCRF4XX_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MCRF4XX), fodder1MB)
}
func Benchmark_SigurnCRC16_RIELLO_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_RIELLO), fodder1MB)
}
func Benchmark_SigurnCRC16_T10_DIF_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_T10_DIF), fodder1MB)
}
func Benchmark_SigurnCRC16_TELEDISK_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TELEDISK), fodder1MB)
}
func Benchmark_SigurnCRC16_TMS37157_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TMS37157), fodder1MB)
}
func Benchmark_SigurnCRC16_USB_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_USB), fodder1MB)
}
func Benchmark_SigurnCRC16_CRC_A_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CRC_A), fodder1MB)
}
func Benchmark_SigurnCRC16_KERMIT_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_KERMIT), fodder1MB)
}
func Benchmark_SigurnCRC16_MODBUS_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MODBUS), fodder1MB)
}
func Benchmark_SigurnCRC16_X_25_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_X_25), fodder1MB)
}
func Benchmark_SigurnCRC16_XMODEM_1MB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_XMODEM), fodder1MB)
}

// 1 KB benchmark
func Benchmark_CCITT_1KB(b *testing.B) {
	checksumTest(b, NewCrc16CCITT(), fodder1KB)
}
func Benchmark_HoweycIBM_1KB(b *testing.B) {
	checksumTest(b, howeyc.NewIBM(), fodder1KB)
}
func Benchmark_EfaultPPP_1KB(b *testing.B) {
	checksumTest(b, efault.New(efault.PPP), fodder1KB)
}
func Benchmark_EfaultModus_1KB(b *testing.B) {
	checksumTest(b, efault.New(efault.Modbus), fodder1KB)
}
func Benchmark_EfaultKermit_1KB(b *testing.B) {
	checksumTest(b, efault.New(efault.Kermit), fodder1KB)
}
func Benchmark_EfaultXModem_1KB(b *testing.B) {
	checksumTest(b, efault.New(efault.XModem), fodder1KB)
}
func Benchmark_JoaoJeronimo_1KB(b *testing.B) {
	checksumTest(b, newJjWrapper(), fodder1KB)
}
func Benchmark_Jiguorui_1KB(b *testing.B) {
	checksumTest(b, newJigWrapper(), fodder1KB)
}
func Benchmark_SigurnCRC16ARC_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_ARC), fodder1KB)
}
func Benchmark_SigurnCRC16_AUG_CCITT_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_AUG_CCITT), fodder1KB)
}
func Benchmark_SigurnCRC16_BUYPASS_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_BUYPASS), fodder1KB)
}
func Benchmark_SigurnCRC16_CCITT_FALSE_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CCITT_FALSE), fodder1KB)
}
func Benchmark_SigurnCRC16_CDMA2000_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CDMA2000), fodder1KB)
}
func Benchmark_SigurnCRC16_DDS_110_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DDS_110), fodder1KB)
}
func Benchmark_SigurnCRC16_DECT_R_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_R), fodder1KB)
}
func Benchmark_SigurnCRC16_DECT_X_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_X), fodder1KB)
}
func Benchmark_SigurnCRC16_DNP_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DNP), fodder1KB)
}
func Benchmark_SigurnCRC16_EN_13757_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_EN_13757), fodder1KB)
}
func Benchmark_SigurnCRC16_GENIBUS_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_GENIBUS), fodder1KB)
}
func Benchmark_SigurnCRC16_MAXIM_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MAXIM), fodder1KB)
}
func Benchmark_SigurnCRC16_MCRF4XX_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MCRF4XX), fodder1KB)
}
func Benchmark_SigurnCRC16_RIELLO_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_RIELLO), fodder1KB)
}
func Benchmark_SigurnCRC16_T10_DIF_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_T10_DIF), fodder1KB)
}
func Benchmark_SigurnCRC16_TELEDISK_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TELEDISK), fodder1KB)
}
func Benchmark_SigurnCRC16_TMS37157_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TMS37157), fodder1KB)
}
func Benchmark_SigurnCRC16_USB_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_USB), fodder1KB)
}
func Benchmark_SigurnCRC16_CRC_A_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CRC_A), fodder1KB)
}
func Benchmark_SigurnCRC16_KERMIT_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_KERMIT), fodder1KB)
}
func Benchmark_SigurnCRC16_MODBUS_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MODBUS), fodder1KB)
}
func Benchmark_SigurnCRC16_X_25_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_X_25), fodder1KB)
}
func Benchmark_SigurnCRC16_XMODEM_1KB(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_XMODEM), fodder1KB)
}

// 64 B benchmark
func Benchmark_CCITTRange_64B(b *testing.B) {
	checksumTest(b, NewCrc16CCITT(), fodder64B)
}
func Benchmark_HoweycIBM_64B(b *testing.B) {
	checksumTest(b, howeyc.NewIBM(), fodder64B)
}
func Benchmark_EfaultPPP_64B(b *testing.B) {
	checksumTest(b, efault.New(efault.PPP), fodder64B)
}
func Benchmark_EfaultModus_64B(b *testing.B) {
	checksumTest(b, efault.New(efault.Modbus), fodder64B)
}
func Benchmark_EfaultKermit_64B(b *testing.B) {
	checksumTest(b, efault.New(efault.Kermit), fodder64B)
}
func Benchmark_EfaultXModem_64B(b *testing.B) {
	checksumTest(b, efault.New(efault.XModem), fodder64B)
}
func Benchmark_JoaoJeronimo_64B(b *testing.B) {
	checksumTest(b, newJjWrapper(), fodder64B)
}
func Benchmark_Jiguorui_64B(b *testing.B) {
	checksumTest(b, newJigWrapper(), fodder64B)
}
func Benchmark_SigurnCRC16ARC_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_ARC), fodder64B)
}
func Benchmark_SigurnCRC16_AUG_CCITT_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_AUG_CCITT), fodder64B)
}
func Benchmark_SigurnCRC16_BUYPASS_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_BUYPASS), fodder64B)
}
func Benchmark_SigurnCRC16_CCITT_FALSE_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CCITT_FALSE), fodder64B)
}
func Benchmark_SigurnCRC16_CDMA2000_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CDMA2000), fodder64B)
}
func Benchmark_SigurnCRC16_DDS_110_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DDS_110), fodder64B)
}
func Benchmark_SigurnCRC16_DECT_R_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_R), fodder64B)
}
func Benchmark_SigurnCRC16_DECT_X_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_X), fodder64B)
}
func Benchmark_SigurnCRC16_DNP_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DNP), fodder64B)
}
func Benchmark_SigurnCRC16_EN_13757_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_EN_13757), fodder64B)
}
func Benchmark_SigurnCRC16_GENIBUS_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_GENIBUS), fodder64B)
}
func Benchmark_SigurnCRC16_MAXIM_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MAXIM), fodder64B)
}
func Benchmark_SigurnCRC16_MCRF4XX_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MCRF4XX), fodder64B)
}
func Benchmark_SigurnCRC16_RIELLO_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_RIELLO), fodder64B)
}
func Benchmark_SigurnCRC16_T10_DIF_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_T10_DIF), fodder64B)
}
func Benchmark_SigurnCRC16_TELEDISK_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TELEDISK), fodder64B)
}
func Benchmark_SigurnCRC16_TMS37157_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TMS37157), fodder64B)
}
func Benchmark_SigurnCRC16_USB_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_USB), fodder64B)
}
func Benchmark_SigurnCRC16_CRC_A_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CRC_A), fodder64B)
}
func Benchmark_SigurnCRC16_KERMIT_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_KERMIT), fodder64B)
}
func Benchmark_SigurnCRC16_MODBUS_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MODBUS), fodder64B)
}
func Benchmark_SigurnCRC16_X_25_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_X_25), fodder64B)
}
func Benchmark_SigurnCRC16_XMODEM_64B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_XMODEM), fodder64B)
}

// 24 B benchmark
func Benchmark_CCITTRange_24B(b *testing.B) {
	checksumTest(b, NewCrc16CCITT(), fodder24B)
}
func Benchmark_HoweycIBM_24B(b *testing.B) {
	checksumTest(b, howeyc.NewIBM(), fodder24B)
}
func Benchmark_EfaultPPP_24B(b *testing.B) {
	checksumTest(b, efault.New(efault.PPP), fodder24B)
}
func Benchmark_EfaultModus_24B(b *testing.B) {
	checksumTest(b, efault.New(efault.Modbus), fodder24B)
}
func Benchmark_EfaultKermit_24B(b *testing.B) {
	checksumTest(b, efault.New(efault.Kermit), fodder24B)
}
func Benchmark_EfaultXModem_24B(b *testing.B) {
	checksumTest(b, efault.New(efault.XModem), fodder24B)
}
func Benchmark_JoaoJeronimo_24B(b *testing.B) {
	checksumTest(b, newJjWrapper(), fodder24B)
}
func Benchmark_Jiguorui_24B(b *testing.B) {
	checksumTest(b, newJigWrapper(), fodder24B)
}
func Benchmark_SigurnCRC16ARC_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_ARC), fodder24B)
}
func Benchmark_SigurnCRC16_AUG_CCITT_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_AUG_CCITT), fodder24B)
}
func Benchmark_SigurnCRC16_BUYPASS_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_BUYPASS), fodder24B)
}
func Benchmark_SigurnCRC16_CCITT_FALSE_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CCITT_FALSE), fodder24B)
}
func Benchmark_SigurnCRC16_CDMA2000_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CDMA2000), fodder24B)
}
func Benchmark_SigurnCRC16_DDS_110_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DDS_110), fodder24B)
}
func Benchmark_SigurnCRC16_DECT_R_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_R), fodder24B)
}
func Benchmark_SigurnCRC16_DECT_X_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DECT_X), fodder24B)
}
func Benchmark_SigurnCRC16_DNP_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_DNP), fodder24B)
}
func Benchmark_SigurnCRC16_EN_13757_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_EN_13757), fodder24B)
}
func Benchmark_SigurnCRC16_GENIBUS_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_GENIBUS), fodder24B)
}
func Benchmark_SigurnCRC16_MAXIM_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MAXIM), fodder24B)
}
func Benchmark_SigurnCRC16_MCRF4XX_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MCRF4XX), fodder24B)
}
func Benchmark_SigurnCRC16_RIELLO_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_RIELLO), fodder24B)
}
func Benchmark_SigurnCRC16_T10_DIF_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_T10_DIF), fodder24B)
}
func Benchmark_SigurnCRC16_TELEDISK_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TELEDISK), fodder24B)
}
func Benchmark_SigurnCRC16_TMS37157_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_TMS37157), fodder24B)
}
func Benchmark_SigurnCRC16_USB_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_USB), fodder24B)
}
func Benchmark_SigurnCRC16_CRC_A_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_CRC_A), fodder24B)
}
func Benchmark_SigurnCRC16_KERMIT_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_KERMIT), fodder24B)
}
func Benchmark_SigurnCRC16_MODBUS_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_MODBUS), fodder24B)
}
func Benchmark_SigurnCRC16_X_25_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_X_25), fodder24B)
}
func Benchmark_SigurnCRC16_XMODEM_24B(b *testing.B) {
	checksumTest(b, newsgWrapper(crc16.CRC16_XMODEM), fodder24B)
}
