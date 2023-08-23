resource "snowmirror_setting" "my_setting" {
  sync = {
    id    = 6
    name  = "Kelvin Sporer"
    table = "...my_table..."
  }
}