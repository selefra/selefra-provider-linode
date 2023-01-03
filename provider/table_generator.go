package provider

import (
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-linode/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableLinodeRegionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeEventGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeImageGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeTokenGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeAccountGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeKubernetesClusterGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeDomainRecordGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeVolumeGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeTypeGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeDomainGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeBucketGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeInstanceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeProfileGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableLinodeTagGenerator{}),
	}
}
