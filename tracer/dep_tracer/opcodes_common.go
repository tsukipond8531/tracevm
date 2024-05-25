package dep_tracer

var OpcodeToString = map[uint8]string{
    OPInitZero:    "INITZERO",
    OPInitCode:    "INITCODE",
    OPCallData:    "CALLDATA",
    OPConstant:    "CONSTANT",
    OPCoinbase:    "COINBASE",
    OPAddress:     "ADDRESS",
    OPOrigin:      "ORIGIN",
    OPCaller:      "CALLER",
    OPCallValue:   "CALLVALUE",
    OPGasPrice:    "GASPRICE",
    OPTimestamp:   "TIMESTAMP",
    OPNumber:      "NUMBER",
    OPDifficulty:  "DIFFICULTY",
    OPRandom:      "RANDOM",
    OPGasLimit:    "GASLIMIT",
    OPPc:          "PC",
    OPMsize:       "MSIZE",
    OPGas:         "GAS",
    OPChainID:     "CHAINID",
    OPBaseFee:     "BASEFEE",
    OPCreateAddr:  "CREATEADDR",
    OPCreate2Addr: "CREATE2ADDR",
    OPCallResult:  "CALLRESULT",
    OPBlobBaseFee: "BLOBBASEFEE",

	OPSlice:           "SLICE",
	OPConcat:          "CONCAT",
    OPSize:            "SIZE",
    OPCodeSize:        "CODESIZE",
    OPAdd:             "ADD",
    OPMul:             "MUL",
    OPSub:             "SUB",
    OPDiv:             "DIV",
    OPSDiv:            "SDIV",
    OPMod:             "MOD",
    OPSMod:            "SMOD",
    OPExp:             "EXP",
    OPSignExtend:      "SIGNEXTEND",
    OPNot:             "NOT",
    OPLt:              "LT",
    OPGt:              "GT",
    OPSlt:             "SLT",
    OPSgt:             "SGT",
    OPEq:              "EQ",
    OPOr:              "OR",
    OPXor:             "XOR",
    OPAddMod:          "ADDMOD",
    OPMulMod:          "MULMOD",
    OPShl:             "SHL",
    OPShr:             "SHR",
    OPSar:             "SAR",
    OPAnd:             "AND",
    OPIsZero:          "ISZERO",
    OPKeccak:          "KECCAK",
    OPCodeKeccak:      "CODEKECCAK",
    OPBalance:         "BALANCE",
    OPBlockHash:       "BLOCKHASH",
    OPEcRecover:       "ECRECOVER",
    OPSha256:          "SHA256",
    OPRipemd160:       "RIPEMD160",
    OPModExp:          "MODEXP",
    OPEcAddX:          "ECADDX",
    OPEcAddY:          "ECADDY",
    OPEcMulX:          "ECMULX",
    OPEcMulY:          "ECMULY",
    OPEcPairing:       "ECPAIRING",
    OPBlake2F:         "BLAKE2F",
    OPBlobHash:        "BLOBHASH",
    OPPointEvaluation: "POINTEVALUATION",

    OPSLoad:  "SLOAD",
    OPSStore: "SSTORE",
}

func OpcodeIsConstant(opcode uint8) bool {
    return opcode < 0xA0
}

func OpcodeIsAddressable(opcode uint8) bool {
    return opcode >= 0xE0
}