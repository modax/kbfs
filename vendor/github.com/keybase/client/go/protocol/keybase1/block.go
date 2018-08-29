// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/block.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type BlockStatus int

const (
	BlockStatus_UNKNOWN  BlockStatus = 0
	BlockStatus_LIVE     BlockStatus = 1
	BlockStatus_ARCHIVED BlockStatus = 2
)

func (o BlockStatus) DeepCopy() BlockStatus { return o }

var BlockStatusMap = map[string]BlockStatus{
	"UNKNOWN":  0,
	"LIVE":     1,
	"ARCHIVED": 2,
}

var BlockStatusRevMap = map[BlockStatus]string{
	0: "UNKNOWN",
	1: "LIVE",
	2: "ARCHIVED",
}

func (e BlockStatus) String() string {
	if v, ok := BlockStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type GetBlockRes struct {
	BlockKey string      `codec:"blockKey" json:"blockKey"`
	Buf      []byte      `codec:"buf" json:"buf"`
	Size     int         `codec:"size" json:"size"`
	Status   BlockStatus `codec:"status" json:"status"`
}

func (o GetBlockRes) DeepCopy() GetBlockRes {
	return GetBlockRes{
		BlockKey: o.BlockKey,
		Buf: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.Buf),
		Size:   o.Size,
		Status: o.Status.DeepCopy(),
	}
}

type BlockRefNonce [8]byte

func (o BlockRefNonce) DeepCopy() BlockRefNonce {
	var ret BlockRefNonce
	copy(ret[:], o[:])
	return ret
}

type BlockReference struct {
	Bid       BlockIdCombo  `codec:"bid" json:"bid"`
	Nonce     BlockRefNonce `codec:"nonce" json:"nonce"`
	ChargedTo UserOrTeamID  `codec:"chargedTo" json:"chargedTo"`
}

func (o BlockReference) DeepCopy() BlockReference {
	return BlockReference{
		Bid:       o.Bid.DeepCopy(),
		Nonce:     o.Nonce.DeepCopy(),
		ChargedTo: o.ChargedTo.DeepCopy(),
	}
}

type BlockReferenceCount struct {
	Ref       BlockReference `codec:"ref" json:"ref"`
	LiveCount int            `codec:"liveCount" json:"liveCount"`
}

func (o BlockReferenceCount) DeepCopy() BlockReferenceCount {
	return BlockReferenceCount{
		Ref:       o.Ref.DeepCopy(),
		LiveCount: o.LiveCount,
	}
}

type DowngradeReferenceRes struct {
	Completed []BlockReferenceCount `codec:"completed" json:"completed"`
	Failed    BlockReference        `codec:"failed" json:"failed"`
}

func (o DowngradeReferenceRes) DeepCopy() DowngradeReferenceRes {
	return DowngradeReferenceRes{
		Completed: (func(x []BlockReferenceCount) []BlockReferenceCount {
			if x == nil {
				return nil
			}
			ret := make([]BlockReferenceCount, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Completed),
		Failed: o.Failed.DeepCopy(),
	}
}

type BlockPingResponse struct {
}

func (o BlockPingResponse) DeepCopy() BlockPingResponse {
	return BlockPingResponse{}
}

type GetSessionChallengeArg struct {
}

type AuthenticateSessionArg struct {
	Signature string `codec:"signature" json:"signature"`
}

type PutBlockArg struct {
	Bid      BlockIdCombo `codec:"bid" json:"bid"`
	Folder   string       `codec:"folder" json:"folder"`
	BlockKey string       `codec:"blockKey" json:"blockKey"`
	Buf      []byte       `codec:"buf" json:"buf"`
}

type PutBlockAgainArg struct {
	Folder   string         `codec:"folder" json:"folder"`
	Ref      BlockReference `codec:"ref" json:"ref"`
	BlockKey string         `codec:"blockKey" json:"blockKey"`
	Buf      []byte         `codec:"buf" json:"buf"`
}

type GetBlockArg struct {
	Bid      BlockIdCombo `codec:"bid" json:"bid"`
	Folder   string       `codec:"folder" json:"folder"`
	SizeOnly bool         `codec:"sizeOnly" json:"sizeOnly"`
}

type AddReferenceArg struct {
	Folder string         `codec:"folder" json:"folder"`
	Ref    BlockReference `codec:"ref" json:"ref"`
}

type DelReferenceArg struct {
	Folder string         `codec:"folder" json:"folder"`
	Ref    BlockReference `codec:"ref" json:"ref"`
}

type ArchiveReferenceArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

type DelReferenceWithCountArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

type ArchiveReferenceWithCountArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

type GetUserQuotaInfoArg struct {
}

type GetTeamQuotaInfoArg struct {
	Tid TeamID `codec:"tid" json:"tid"`
}

type BlockPingArg struct {
}

type BlockInterface interface {
	GetSessionChallenge(context.Context) (ChallengeInfo, error)
	AuthenticateSession(context.Context, string) error
	PutBlock(context.Context, PutBlockArg) error
	PutBlockAgain(context.Context, PutBlockAgainArg) error
	GetBlock(context.Context, GetBlockArg) (GetBlockRes, error)
	AddReference(context.Context, AddReferenceArg) error
	DelReference(context.Context, DelReferenceArg) error
	ArchiveReference(context.Context, ArchiveReferenceArg) ([]BlockReference, error)
	DelReferenceWithCount(context.Context, DelReferenceWithCountArg) (DowngradeReferenceRes, error)
	ArchiveReferenceWithCount(context.Context, ArchiveReferenceWithCountArg) (DowngradeReferenceRes, error)
	GetUserQuotaInfo(context.Context) ([]byte, error)
	GetTeamQuotaInfo(context.Context, TeamID) ([]byte, error)
	BlockPing(context.Context) (BlockPingResponse, error)
}

func BlockProtocol(i BlockInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.block",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getSessionChallenge": {
				MakeArg: func() interface{} {
					ret := make([]GetSessionChallengeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetSessionChallenge(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"authenticateSession": {
				MakeArg: func() interface{} {
					ret := make([]AuthenticateSessionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AuthenticateSessionArg)
					if !ok {
						err = rpc.NewTypeError((*[]AuthenticateSessionArg)(nil), args)
						return
					}
					err = i.AuthenticateSession(ctx, (*typedArgs)[0].Signature)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"putBlock": {
				MakeArg: func() interface{} {
					ret := make([]PutBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutBlockArg)(nil), args)
						return
					}
					err = i.PutBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"putBlockAgain": {
				MakeArg: func() interface{} {
					ret := make([]PutBlockAgainArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutBlockAgainArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutBlockAgainArg)(nil), args)
						return
					}
					err = i.PutBlockAgain(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getBlock": {
				MakeArg: func() interface{} {
					ret := make([]GetBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetBlockArg)(nil), args)
						return
					}
					ret, err = i.GetBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"addReference": {
				MakeArg: func() interface{} {
					ret := make([]AddReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AddReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]AddReferenceArg)(nil), args)
						return
					}
					err = i.AddReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delReference": {
				MakeArg: func() interface{} {
					ret := make([]DelReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DelReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]DelReferenceArg)(nil), args)
						return
					}
					err = i.DelReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"archiveReference": {
				MakeArg: func() interface{} {
					ret := make([]ArchiveReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ArchiveReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]ArchiveReferenceArg)(nil), args)
						return
					}
					ret, err = i.ArchiveReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delReferenceWithCount": {
				MakeArg: func() interface{} {
					ret := make([]DelReferenceWithCountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DelReferenceWithCountArg)
					if !ok {
						err = rpc.NewTypeError((*[]DelReferenceWithCountArg)(nil), args)
						return
					}
					ret, err = i.DelReferenceWithCount(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"archiveReferenceWithCount": {
				MakeArg: func() interface{} {
					ret := make([]ArchiveReferenceWithCountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ArchiveReferenceWithCountArg)
					if !ok {
						err = rpc.NewTypeError((*[]ArchiveReferenceWithCountArg)(nil), args)
						return
					}
					ret, err = i.ArchiveReferenceWithCount(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getUserQuotaInfo": {
				MakeArg: func() interface{} {
					ret := make([]GetUserQuotaInfoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetUserQuotaInfo(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getTeamQuotaInfo": {
				MakeArg: func() interface{} {
					ret := make([]GetTeamQuotaInfoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetTeamQuotaInfoArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetTeamQuotaInfoArg)(nil), args)
						return
					}
					ret, err = i.GetTeamQuotaInfo(ctx, (*typedArgs)[0].Tid)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"blockPing": {
				MakeArg: func() interface{} {
					ret := make([]BlockPingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.BlockPing(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type BlockClient struct {
	Cli rpc.GenericClient
}

func (c BlockClient) GetSessionChallenge(ctx context.Context) (res ChallengeInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getSessionChallenge", []interface{}{GetSessionChallengeArg{}}, &res)
	return
}

func (c BlockClient) AuthenticateSession(ctx context.Context, signature string) (err error) {
	__arg := AuthenticateSessionArg{Signature: signature}
	err = c.Cli.Call(ctx, "keybase.1.block.authenticateSession", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) PutBlock(ctx context.Context, __arg PutBlockArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.putBlock", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) PutBlockAgain(ctx context.Context, __arg PutBlockAgainArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.putBlockAgain", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) GetBlock(ctx context.Context, __arg GetBlockArg) (res GetBlockRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getBlock", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) AddReference(ctx context.Context, __arg AddReferenceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.addReference", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) DelReference(ctx context.Context, __arg DelReferenceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.delReference", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) ArchiveReference(ctx context.Context, __arg ArchiveReferenceArg) (res []BlockReference, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.archiveReference", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) DelReferenceWithCount(ctx context.Context, __arg DelReferenceWithCountArg) (res DowngradeReferenceRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.delReferenceWithCount", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) ArchiveReferenceWithCount(ctx context.Context, __arg ArchiveReferenceWithCountArg) (res DowngradeReferenceRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.archiveReferenceWithCount", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) GetUserQuotaInfo(ctx context.Context) (res []byte, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getUserQuotaInfo", []interface{}{GetUserQuotaInfoArg{}}, &res)
	return
}

func (c BlockClient) GetTeamQuotaInfo(ctx context.Context, tid TeamID) (res []byte, err error) {
	__arg := GetTeamQuotaInfoArg{Tid: tid}
	err = c.Cli.Call(ctx, "keybase.1.block.getTeamQuotaInfo", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) BlockPing(ctx context.Context) (res BlockPingResponse, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.blockPing", []interface{}{BlockPingArg{}}, &res)
	return
}
