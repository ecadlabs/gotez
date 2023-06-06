package expression

import "strconv"

type Prim uint8

func (prim Prim) String() string {
	if s, ok := primStr[prim]; ok {
		return s
	}
	return strconv.FormatInt(int64(prim), 10)
}

func (prim Prim) MarshalText() (text []byte, err error) {
	return []byte(prim.String()), nil
}

const (
	Prim_parameter                      Prim = 0
	Prim_storage                        Prim = 1
	Prim_code                           Prim = 2
	Prim_False                          Prim = 3
	Prim_Elt                            Prim = 4
	Prim_Left                           Prim = 5
	Prim_None                           Prim = 6
	Prim_Pair                           Prim = 7
	Prim_Right                          Prim = 8
	Prim_Some                           Prim = 9
	Prim_True                           Prim = 10
	Prim_Unit                           Prim = 11
	Prim_PACK                           Prim = 12
	Prim_UNPACK                         Prim = 13
	Prim_BLAKE2B                        Prim = 14
	Prim_SHA256                         Prim = 15
	Prim_SHA512                         Prim = 16
	Prim_ABS                            Prim = 17
	Prim_ADD                            Prim = 18
	Prim_AMOUNT                         Prim = 19
	Prim_AND                            Prim = 20
	Prim_BALANCE                        Prim = 21
	Prim_CAR                            Prim = 22
	Prim_CDR                            Prim = 23
	Prim_CHECK_SIGNATURE                Prim = 24
	Prim_COMPARE                        Prim = 25
	Prim_CONCAT                         Prim = 26
	Prim_CONS                           Prim = 27
	Prim_CREATE_ACCOUNT                 Prim = 28
	Prim_CREATE_CONTRACT                Prim = 29
	Prim_IMPLICIT_ACCOUNT               Prim = 30
	Prim_DIP                            Prim = 31
	Prim_DROP                           Prim = 32
	Prim_DUP                            Prim = 33
	Prim_EDIV                           Prim = 34
	Prim_EMPTY_MAP                      Prim = 35
	Prim_EMPTY_SET                      Prim = 36
	Prim_EQ                             Prim = 37
	Prim_EXEC                           Prim = 38
	Prim_FAILWITH                       Prim = 39
	Prim_GE                             Prim = 40
	Prim_GET                            Prim = 41
	Prim_GT                             Prim = 42
	Prim_HASH_KEY                       Prim = 43
	Prim_IF                             Prim = 44
	Prim_IF_CONS                        Prim = 45
	Prim_IF_LEFT                        Prim = 46
	Prim_IF_NONE                        Prim = 47
	Prim_INT                            Prim = 48
	Prim_LAMBDA                         Prim = 49
	Prim_LE                             Prim = 50
	Prim_LEFT                           Prim = 51
	Prim_LOOP                           Prim = 52
	Prim_LSL                            Prim = 53
	Prim_LSR                            Prim = 54
	Prim_LT                             Prim = 55
	Prim_MAP                            Prim = 56
	Prim_MEM                            Prim = 57
	Prim_MUL                            Prim = 58
	Prim_NEG                            Prim = 59
	Prim_NEQ                            Prim = 60
	Prim_NIL                            Prim = 61
	Prim_NONE                           Prim = 62
	Prim_NOT                            Prim = 63
	Prim_NOW                            Prim = 64
	Prim_OR                             Prim = 65
	Prim_PAIR                           Prim = 66
	Prim_PUSH                           Prim = 67
	Prim_RIGHT                          Prim = 68
	Prim_SIZE                           Prim = 69
	Prim_SOME                           Prim = 70
	Prim_SOURCE                         Prim = 71
	Prim_SENDER                         Prim = 72
	Prim_SELF                           Prim = 73
	Prim_STEPS_TO_QUOTA                 Prim = 74
	Prim_SUB                            Prim = 75
	Prim_SWAP                           Prim = 76
	Prim_TRANSFER_TOKENS                Prim = 77
	Prim_SET_DELEGATE                   Prim = 78
	Prim_UNIT                           Prim = 79
	Prim_UPDATE                         Prim = 80
	Prim_XOR                            Prim = 81
	Prim_ITER                           Prim = 82
	Prim_LOOP_LEFT                      Prim = 83
	Prim_ADDRESS                        Prim = 84
	Prim_CONTRACT                       Prim = 85
	Prim_ISNAT                          Prim = 86
	Prim_CAST                           Prim = 87
	Prim_RENAME                         Prim = 88
	Prim_bool                           Prim = 89
	Prim_contract                       Prim = 90
	Prim_int                            Prim = 91
	Prim_key                            Prim = 92
	Prim_key_hash                       Prim = 93
	Prim_lambda                         Prim = 94
	Prim_list                           Prim = 95
	Prim_map                            Prim = 96
	Prim_big_map                        Prim = 97
	Prim_nat                            Prim = 98
	Prim_option                         Prim = 99
	Prim_or                             Prim = 100
	Prim_pair                           Prim = 101
	Prim_set                            Prim = 102
	Prim_signature                      Prim = 103
	Prim_string                         Prim = 104
	Prim_bytes                          Prim = 105
	Prim_mutez                          Prim = 106
	Prim_timestamp                      Prim = 107
	Prim_unit                           Prim = 108
	Prim_operation                      Prim = 109
	Prim_address                        Prim = 110
	Prim_SLICE                          Prim = 111
	Prim_DIG                            Prim = 112
	Prim_DUG                            Prim = 113
	Prim_EMPTY_BIG_MAP                  Prim = 114
	Prim_APPLY                          Prim = 115
	Prim_chain_id                       Prim = 116
	Prim_CHAIN_ID                       Prim = 117
	Prim_LEVEL                          Prim = 118
	Prim_SELF_ADDRESS                   Prim = 119
	Prim_never                          Prim = 120
	Prim_NEVER                          Prim = 121
	Prim_UNPAIR                         Prim = 122
	Prim_VOTING_POWER                   Prim = 123
	Prim_TOTAL_VOTING_POWER             Prim = 124
	Prim_KECCAK                         Prim = 125
	Prim_SHA3                           Prim = 126
	Prim_PAIRING_CHECK                  Prim = 127
	Prim_bls12_381_g1                   Prim = 128
	Prim_bls12_381_g2                   Prim = 129
	Prim_bls12_381_fr                   Prim = 130
	Prim_sapling_state                  Prim = 131
	Prim_sapling_transaction_deprecated Prim = 132
	Prim_SAPLING_EMPTY_STATE            Prim = 133
	Prim_SAPLING_VERIFY_UPDATE          Prim = 134
	Prim_ticket                         Prim = 135
	Prim_TICKET_DEPRECATED              Prim = 136
	Prim_READ_TICKET                    Prim = 137
	Prim_SPLIT_TICKET                   Prim = 138
	Prim_JOIN_TICKETS                   Prim = 139
	Prim_GET_AND_UPDATE                 Prim = 140
	Prim_chest                          Prim = 141
	Prim_chest_key                      Prim = 142
	Prim_OPEN_CHEST                     Prim = 143
	Prim_VIEW                           Prim = 144
	Prim_view                           Prim = 145
	Prim_constant                       Prim = 146
	Prim_SUB_MUTEZ                      Prim = 147
	Prim_tx_rollup_l2_address           Prim = 148
	Prim_MIN_BLOCK_TIME                 Prim = 149
	Prim_sapling_transaction            Prim = 150
	Prim_EMIT                           Prim = 151
	Prim_Lambda_rec                     Prim = 152
	Prim_LAMBDA_REC                     Prim = 153
	Prim_TICKET                         Prim = 154
	Prim_BYTES                          Prim = 155
	Prim_NAT                            Prim = 156
)

var primStr = map[Prim]string{
	Prim_parameter:                      "parameter",
	Prim_storage:                        "storage",
	Prim_code:                           "code",
	Prim_False:                          "False",
	Prim_Elt:                            "Elt",
	Prim_Left:                           "Left",
	Prim_None:                           "None",
	Prim_Pair:                           "Pair",
	Prim_Right:                          "Right",
	Prim_Some:                           "Some",
	Prim_True:                           "True",
	Prim_Unit:                           "Unit",
	Prim_PACK:                           "PACK",
	Prim_UNPACK:                         "UNPACK",
	Prim_BLAKE2B:                        "BLAKE2B",
	Prim_SHA256:                         "SHA256",
	Prim_SHA512:                         "SHA512",
	Prim_ABS:                            "ABS",
	Prim_ADD:                            "ADD",
	Prim_AMOUNT:                         "AMOUNT",
	Prim_AND:                            "AND",
	Prim_BALANCE:                        "BALANCE",
	Prim_CAR:                            "CAR",
	Prim_CDR:                            "CDR",
	Prim_CHECK_SIGNATURE:                "CHECK_SIGNATURE",
	Prim_COMPARE:                        "COMPARE",
	Prim_CONCAT:                         "CONCAT",
	Prim_CONS:                           "CONS",
	Prim_CREATE_ACCOUNT:                 "CREATE_ACCOUNT",
	Prim_CREATE_CONTRACT:                "CREATE_CONTRACT",
	Prim_IMPLICIT_ACCOUNT:               "IMPLICIT_ACCOUNT",
	Prim_DIP:                            "DIP",
	Prim_DROP:                           "DROP",
	Prim_DUP:                            "DUP",
	Prim_EDIV:                           "EDIV",
	Prim_EMPTY_MAP:                      "EMPTY_MAP",
	Prim_EMPTY_SET:                      "EMPTY_SET",
	Prim_EQ:                             "EQ",
	Prim_EXEC:                           "EXEC",
	Prim_FAILWITH:                       "FAILWITH",
	Prim_GE:                             "GE",
	Prim_GET:                            "GET",
	Prim_GT:                             "GT",
	Prim_HASH_KEY:                       "HASH_KEY",
	Prim_IF:                             "IF",
	Prim_IF_CONS:                        "IF_CONS",
	Prim_IF_LEFT:                        "IF_LEFT",
	Prim_IF_NONE:                        "IF_NONE",
	Prim_INT:                            "INT",
	Prim_LAMBDA:                         "LAMBDA",
	Prim_LE:                             "LE",
	Prim_LEFT:                           "LEFT",
	Prim_LOOP:                           "LOOP",
	Prim_LSL:                            "LSL",
	Prim_LSR:                            "LSR",
	Prim_LT:                             "LT",
	Prim_MAP:                            "MAP",
	Prim_MEM:                            "MEM",
	Prim_MUL:                            "MUL",
	Prim_NEG:                            "NEG",
	Prim_NEQ:                            "NEQ",
	Prim_NIL:                            "NIL",
	Prim_NONE:                           "NONE",
	Prim_NOT:                            "NOT",
	Prim_NOW:                            "NOW",
	Prim_OR:                             "OR",
	Prim_PAIR:                           "PAIR",
	Prim_PUSH:                           "PUSH",
	Prim_RIGHT:                          "RIGHT",
	Prim_SIZE:                           "SIZE",
	Prim_SOME:                           "SOME",
	Prim_SOURCE:                         "SOURCE",
	Prim_SENDER:                         "SENDER",
	Prim_SELF:                           "SELF",
	Prim_STEPS_TO_QUOTA:                 "STEPS_TO_QUOTA",
	Prim_SUB:                            "SUB",
	Prim_SWAP:                           "SWAP",
	Prim_TRANSFER_TOKENS:                "TRANSFER_TOKENS",
	Prim_SET_DELEGATE:                   "SET_DELEGATE",
	Prim_UNIT:                           "UNIT",
	Prim_UPDATE:                         "UPDATE",
	Prim_XOR:                            "XOR",
	Prim_ITER:                           "ITER",
	Prim_LOOP_LEFT:                      "LOOP_LEFT",
	Prim_ADDRESS:                        "ADDRESS",
	Prim_CONTRACT:                       "CONTRACT",
	Prim_ISNAT:                          "ISNAT",
	Prim_CAST:                           "CAST",
	Prim_RENAME:                         "RENAME",
	Prim_bool:                           "bool",
	Prim_contract:                       "contract",
	Prim_int:                            "int",
	Prim_key:                            "key",
	Prim_key_hash:                       "key_hash",
	Prim_lambda:                         "lambda",
	Prim_list:                           "list",
	Prim_map:                            "map",
	Prim_big_map:                        "big_map",
	Prim_nat:                            "nat",
	Prim_option:                         "option",
	Prim_or:                             "or",
	Prim_pair:                           "pair",
	Prim_set:                            "set",
	Prim_signature:                      "signature",
	Prim_string:                         "string",
	Prim_bytes:                          "bytes",
	Prim_mutez:                          "mutez",
	Prim_timestamp:                      "timestamp",
	Prim_unit:                           "unit",
	Prim_operation:                      "operation",
	Prim_address:                        "address",
	Prim_SLICE:                          "SLICE",
	Prim_DIG:                            "DIG",
	Prim_DUG:                            "DUG",
	Prim_EMPTY_BIG_MAP:                  "EMPTY_BIG_MAP",
	Prim_APPLY:                          "APPLY",
	Prim_chain_id:                       "chain_id",
	Prim_CHAIN_ID:                       "CHAIN_ID",
	Prim_LEVEL:                          "LEVEL",
	Prim_SELF_ADDRESS:                   "SELF_ADDRESS",
	Prim_never:                          "never",
	Prim_NEVER:                          "NEVER",
	Prim_UNPAIR:                         "UNPAIR",
	Prim_VOTING_POWER:                   "VOTING_POWER",
	Prim_TOTAL_VOTING_POWER:             "TOTAL_VOTING_POWER",
	Prim_KECCAK:                         "KECCAK",
	Prim_SHA3:                           "SHA3",
	Prim_PAIRING_CHECK:                  "PAIRING_CHECK",
	Prim_bls12_381_g1:                   "bls12_381_g1",
	Prim_bls12_381_g2:                   "bls12_381_g2",
	Prim_bls12_381_fr:                   "bls12_381_fr",
	Prim_sapling_state:                  "sapling_state",
	Prim_sapling_transaction_deprecated: "sapling_transaction_deprecated",
	Prim_SAPLING_EMPTY_STATE:            "SAPLING_EMPTY_STATE",
	Prim_SAPLING_VERIFY_UPDATE:          "SAPLING_VERIFY_UPDATE",
	Prim_ticket:                         "ticket",
	Prim_TICKET_DEPRECATED:              "TICKET_DEPRECATED",
	Prim_READ_TICKET:                    "READ_TICKET",
	Prim_SPLIT_TICKET:                   "SPLIT_TICKET",
	Prim_JOIN_TICKETS:                   "JOIN_TICKETS",
	Prim_GET_AND_UPDATE:                 "GET_AND_UPDATE",
	Prim_chest:                          "chest",
	Prim_chest_key:                      "chest_key",
	Prim_OPEN_CHEST:                     "OPEN_CHEST",
	Prim_VIEW:                           "VIEW",
	Prim_view:                           "view",
	Prim_constant:                       "constant",
	Prim_SUB_MUTEZ:                      "SUB_MUTEZ",
	Prim_tx_rollup_l2_address:           "tx_rollup_l2_address",
	Prim_MIN_BLOCK_TIME:                 "MIN_BLOCK_TIME",
	Prim_sapling_transaction:            "sapling_transaction",
	Prim_EMIT:                           "EMIT",
	Prim_Lambda_rec:                     "Lambda_rec",
	Prim_LAMBDA_REC:                     "LAMBDA_REC",
	Prim_TICKET:                         "TICKET",
	Prim_BYTES:                          "BYTES",
	Prim_NAT:                            "NAT",
}
