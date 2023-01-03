package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeProfileGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeProfileGenerator{}

func (x *TableLinodeProfileGenerator) GetTableName() string {
	return "linode_profile"
}

func (x *TableLinodeProfileGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeProfileGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeProfileGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeProfileGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			item, err := conn.GetProfile(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- item
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableLinodeProfileGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeProfileGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Description("Your username, used for logging in to our system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Your email address. This address will be used for communication with Linode as necessary.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_notifications").ColumnType(schema.ColumnTypeBool).Description("If true, you will receive email notifications about account activity. If false, you may still receive business-critical communications through email.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lish_auth_method").ColumnType(schema.ColumnTypeString).Description("The authentication methods that are allowed when connecting to the Linode Shell (Lish): password_keys, keys_only, disabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("referrals").ColumnType(schema.ColumnTypeJSON).Description("Information about your status in our referral program. This information becomes accessible after this Profile’s Account has established at least $25.00 USD of total payments.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restricted").ColumnType(schema.ColumnTypeBool).Description("If true, your User has restrictions on what can be accessed on your Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("two_factor_auth").ColumnType(schema.ColumnTypeBool).Description("If true, logins from untrusted computers will require Two Factor Authentication.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).Description("Your unique ID in our system. This value will never change, and can safely be used to identify your User.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorized_keys").ColumnType(schema.ColumnTypeJSON).Description("The list of SSH Keys authorized to use Lish for your User.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_whitelist_enabled").ColumnType(schema.ColumnTypeBool).Description("If true, logins for your User will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).Description("The timezone you prefer to see times in.").Build(),
	}
}

func (x *TableLinodeProfileGenerator) GetSubTables() []*schema.Table {
	return nil
}
