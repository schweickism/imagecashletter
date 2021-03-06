package imagecashletter

import (
	"testing"
)

func TestCashLetterPanics(t *testing.T) {
	var cl *CashLetter

	if v := cl.GetBundles(); v != nil {
		t.Errorf("unexpected GetBundles: %v", v)
	}
	if v := cl.GetRoutingNumberSummary(); v != nil {
		t.Errorf("unexpected GetRoutingNumberSummary: %v", v)
	}
	if v := cl.GetCreditItems(); v != nil {
		t.Errorf("unexpected GetCreditItems: %v", v)
	}
}

// TestCashLetterNoBundle validates no Bundle when CashLetterHeader.RecordTypeIndicator = "N"
func TestCashLetterNoBundle(t *testing.T) {
	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.GetHeader().RecordTypeIndicator = "N"
	cl.AddBundle(bundle)
	if err := cl.Create(); err != nil {
		if e, ok := err.(*CashLetterError); ok {
			if e.FieldName != "RecordTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCashLetterNoRoutingNumberSummary validates no Bundle when CashLetterHeader.CollectionTypeIndicator is not
// 00, 01, 02
func TestCashLetterRoutingNumberSummary(t *testing.T) {
	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.GetHeader().CollectionTypeIndicator = "03"
	cl.AddBundle(bundle)
	rns := mockRoutingNumberSummary()
	cl.AddRoutingNumberSummary(rns)
	if err := cl.Create(); err != nil {
		if e, ok := err.(*CashLetterError); ok {
			if e.FieldName != "CollectionTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
