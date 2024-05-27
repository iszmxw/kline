package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"goapi/pkg/logger"
	"google.golang.org/api/option"
)

func SubscribeToTopic(serviceAccountKeyJSON, topic string, tokens []string) (response *messaging.TopicManagementResponse) {
	// 使用服务账号密钥文件初始化 Firebase 应用
	opt := option.WithCredentialsJSON([]byte(serviceAccountKeyJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error(err)
		return
	}

	// 获取 Firebase Cloud Messaging 客户端
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	// 执行订阅操作
	response, err = client.SubscribeToTopic(ctx, tokens, topic)
	if err != nil {
		logger.Error(err)
	}
	return
}

func UnsubscribeFromTopic(serviceAccountKeyJSON, topic string, tokens []string) (response *messaging.TopicManagementResponse) {
	// 使用服务账号密钥文件初始化 Firebase 应用
	opt := option.WithCredentialsJSON([]byte(serviceAccountKeyJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error(err)
		return
	}

	// 获取 Firebase Cloud Messaging 客户端
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	// 执行取消订阅操作
	response, err = client.UnsubscribeFromTopic(ctx, tokens, topic)
	if err != nil {
		logger.Error(err)
	}
	return
}

func TestMain() {
	// 直接提供服务账号密钥的 JSON 字符串
	serviceAccountKeyJSON := `{
		"type": "service_account",
		"project_id": "smartoption-10620",
		"private_key_id": "3b3da597e48bad81ffeccb2969f8885b876208ce",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCP4csj7GPz3DCp\nGrxkfbUYatk2wVJzlSlVD5rMd3OgI16tCA22IjV2XBylFRRj373m+ZFiPjIvwUhd\nN/c31sucpGOlJQqu+p4KuOg2hy1TAnWPyYkt0yLC58aHQxTZa1lPdbX2TXnnT0TE\nIMQ5wYRdo6lvAETYXk71iG6mzZJw8PQ2ihRw/OdPZXC2hGoGgdWidC2noMGhmYj4\n3MeNvwca/UX+d+AbHNlVnEYUjQ5L+udqMmlo4k/4sQilAK5nWy5VTOwAuhrpXIpz\no4rE/rC7tEr8oNdzH6geh/LzLdIV7KwQZB9naG+M1AnJ9ac3nTKoSsvAZRRL6p9u\nqlqnYtkfAgMBAAECggEAKsIbMG3Ltmi1PUvouKRy8PKv4onJCvhE9DZS9nh5MBnS\nxPEjwgCgirANAzIkgRm4wa9g+SUqnO5/ptQ+76UxlTd1NVqwthdcaCtTNwEtAwI0\nFar6zVd1eR6P7RqOU56HSAnfLXcQxyDTesWSCufrgjhk3cviqbo9FNKar49f+Zfv\noMlQd0MynpNdTsMIKZqlCVqRSQ359BvY2HL+zVS/hks8S2xOr/P3y7yPfb8TuLnC\nSQ8LyDWW0GupLUywl9OwvPtQ0uDX0AmzdUnZu9JRYcUXW0gCBdbXmmQGlybkZ4YU\nk23mFqWOTU7tEuBvgbJl2EM+axpYrI0QnAK7zKD34QKBgQDEnRYwXyH63wao0F+t\nYqScvciZQ0WBRf/r42g4ailgolwt77v3zByzkaUXhrzqGCBaM8h25xnTLxUTUj7j\nAaISUc4BqsYmTyEkPkCdJA75+6RyfH6cFO5a4FpNOoB4QHs+dYWm38mQqN20yRXv\nLSURAtjpui0S+lbG7v/spz3oHQKBgQC7V0w7IWwrqpmd3JHm53EkBR2jYa9IiFHi\no9Zc+cqVqLORNrzTfAZNaa9sdcBZTvbZ8E99X79aD1jZLoGteVStzWSM4GuWKWF4\n1ZKudghN6NRdnzcj/832vJRe9n0kpWg5y9a40y3efBr/9jvANweskt7kujvzjs63\nn6Q6YLgZawKBgBiUTs7TRcs4EHwQOLXTTCLwMuC+ITVNC/ZwYnoTA4M6KO3xN319\n/eSRA6OSS5G0XG8KZDhBIeBd1XOp54S3koxZv7hDoivt3oXKZod+l4S26dMKL0XR\n9A5WV9OpyZ++5+liBchb17YG9liAfwZ8jVK4GWf2x2XE+A48mK/8JGgdAoGAEKma\nESRO369jzGOi4nuxV3fRBAgbNANT7SmPVmOtPhRsLD1HHDRKxzXgMXBjtE3OrOGv\n4EHU3ijBSUfWptxEyNSc3d7UEt4JWlr8meHLqUpNISqZF+2Xnxt0wuF/jfPcSk+0\nAGpgCQrWx+9ZryX128aQZp2d2XLZNgHP9JVHqpMCgYB/G8V3QOBjjtADldQQFskR\nRmp0F9ivIhXu1O+cwHGHaxHNK9WQAMRQ+bDI+sJA0IGtZzwHcyGDsR/7T7Qk8noI\nN4zlhABdzdtpuJseFZRxF+pFZ7aO7M9M1RTSrQgeLHe4BBFouE4Y10dYoMOOqlzb\nlzmliQonDN9a30jgRKcIqg==\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-2whzw@smartoption-10620.iam.gserviceaccount.com",
		"client_id": "116600938784560506381",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-2whzw%40smartoption-10620.iam.gserviceaccount.com"
	}`
	// 要订阅的主题
	topic := "smartOptionIndonesia"
	// 要订阅的设备令牌
	tokens := []string{
		"eria9ceXQViaeGX5juiWe5:APA91bHrMCanre4iJP9Q7IfgqQAkNjR6Ah0vWy2cA7UxJUH7giVIFUAsiPVznrDTgpp6qRhuTNDHpn5O8g0f4GfyanaM84guex45DsCDAm7_QRfVcNCRRBWTEyL1GOq4uurV8tWlKWCb",
	}
	response := SubscribeToTopic(serviceAccountKeyJSON, topic, tokens)
	fmt.Printf("Successfully subscribed to topic: %d tokens were subscribed\n", response.SuccessCount)
}
