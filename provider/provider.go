package provider

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
	"os"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "linode",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var linodeConfig linode_client.Configs

				err := config.Unmarshal(&linodeConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(linodeConfig.Providers) == 0 {
					linodeConfig.Providers = append(linodeConfig.Providers, linode_client.Config{})
				}

				if linodeConfig.Providers[0].Token == "" {
					linodeConfig.Providers[0].Token = os.Getenv("LINODE_TOKEN")
				}

				if linodeConfig.Providers[0].Token == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing token in configuration")
				}

				clients, err := linode_client.NewClients(linodeConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# token: "<YOUR_ACCESS_TOKEN>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var clientConfig linode_client.Configs
				err := config.Unmarshal(&clientConfig.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
