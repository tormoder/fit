package profile

type datarowpos int

// Row indexes for raw type data.
const (
	tNAME datarowpos = iota
	tBTYPE
	tVALNAME
	tVAL
	tCOMMENT
)

// Row indexes for raw message data.
const (
	mMSGNAME datarowpos = iota
	mFDEFN
	mFNAME
	mFTYPE
	mARRAY
	mCOMPS
	mSCALE
	mOFFSET
	mUNITS
	mBITS
	mACCUMU
	mRFNAME
	mRFVAL
	mCOMMENT
	mPRODS
	mEXAMPLE
)
