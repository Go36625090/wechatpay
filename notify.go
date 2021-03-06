package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
)

type NotifyService struct {
	appID string
	mchID string
	notifyURL string
	*notify.Handler
}

var certificateVisitor core.CertificateVisitor

func (m *Client) GetNotifyService() (*NotifyService, error) {
	ctx := context.Background()
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, m.privateKey, m.mchCertificateSerialNumber, m.mchID, m.mchAPIv3Key)
	if err != nil {
		return nil, err
	}

	// 2. 获取商户号对应的微信支付平台证书访问器
	if nil == certificateVisitor{
		certificateVisitor = downloader.MgrInstance().GetCertificateVisitor(m.mchID)
	}
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(m.mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	return &NotifyService{
		appID:     "",
		mchID:     "",
		notifyURL: "",
		Handler:   handler,
	}, nil
}
