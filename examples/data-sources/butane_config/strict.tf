data "butane_config" "config" {
  file   = file("./config.bu")
  strict = true
}
