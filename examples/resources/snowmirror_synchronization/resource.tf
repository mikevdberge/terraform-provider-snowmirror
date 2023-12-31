resource "snowmirror_synchronization" "my_synchronization" {
  active                  = true
  allow_inherited_columns = true
  auto_schema_update      = true
  columns = [
    {
      name = "sys_id"
    },
  ]
  columns_to_exclude = [
    {
      name = "sys_created_by"
    },
  ]
  delete_strategy = "AUDIT"
  encoded_query   = ""
  full_load_scheduler = {
    begin_date     = "2014-08-07"
    execution_type = "CLEAN_AND_SYNCHRONIZE"
    type           = "MANUALLY"
  }
  mirror_table         = "incident"
  name                 = "incident"
  reference_field_type = "...my_reference_field_type..."
  run_immediately      = false
  scheduler = {
    begin_date = "2014-08-07"
    type       = "MANUALLY"
  }
  scheduler_priority = "NORMAL"
  table              = "incident"
  view               = "...my_view..."
}