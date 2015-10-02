package profile

type csvrowpos int

// Row indexes for type csv file.
const (
	tNAME csvrowpos = iota
	tBTYPE
	tVALNAME
	tVAL
	tCOMMENT
)

// Row indexes for message csv file.
const (
	mMSGNAME csvrowpos = iota
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
