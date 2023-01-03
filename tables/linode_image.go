package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeImageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeImageGenerator{}

func (x *TableLinodeImageGenerator) GetTableName() string {
	return "linode_image"
}

func (x *TableLinodeImageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeImageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeImageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeImageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			items, err := conn.ListImages(ctx, &opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range items {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableLinodeImageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeImageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("image_type").ColumnType(schema.ColumnTypeString).Description("How the Image was created: manual, automatic.").
			Extractor(column_value_extractor.StructSelector("type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_public").ColumnType(schema.ColumnTypeBool).Description("True if the Image is public.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("The minimum size this Image needs to deploy. Size is in MB.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique ID of this Image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("A short description of the Image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("When this Image was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deprecated").ColumnType(schema.ColumnTypeBool).Description("Whether or not this Image is deprecated. Will only be true for deprecated public Images.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("The name of the User who created this Image, or 'linode' for official Images.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A detailed description of this Image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiry").ColumnType(schema.ColumnTypeTimestamp).Description("Only Images created automatically (from a deleted Linode; type=automatic) will expire.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vendor").ColumnType(schema.ColumnTypeString).Description("The upstream distribution vendor. None for private Images.").Build(),
	}
}

func (x *TableLinodeImageGenerator) GetSubTables() []*schema.Table {
	return nil
}
