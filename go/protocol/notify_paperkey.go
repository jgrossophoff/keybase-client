// Auto-generated by avdl-compiler v1.3.4 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/notify_paperkey.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type PaperKeyCachedArg struct {
	Uid    UID `codec:"uid" json:"uid"`
	EncKID KID `codec:"encKID" json:"encKID"`
	SigKID KID `codec:"sigKID" json:"sigKID"`
}

type NotifyPaperKeyInterface interface {
	PaperKeyCached(context.Context, PaperKeyCachedArg) error
}

func NotifyPaperKeyProtocol(i NotifyPaperKeyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyPaperKey",
		Methods: map[string]rpc.ServeHandlerDescription{
			"paperKeyCached": {
				MakeArg: func() interface{} {
					ret := make([]PaperKeyCachedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PaperKeyCachedArg)
					if !ok {
						err = rpc.NewTypeError((*[]PaperKeyCachedArg)(nil), args)
						return
					}
					err = i.PaperKeyCached(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyPaperKeyClient struct {
	Cli rpc.GenericClient
}

func (c NotifyPaperKeyClient) PaperKeyCached(ctx context.Context, __arg PaperKeyCachedArg) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyPaperKey.paperKeyCached", []interface{}{__arg})
	return
}
