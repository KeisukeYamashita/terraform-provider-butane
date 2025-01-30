data "butane_config" "config" {
  content = file("./config.bu")
  strict  = true
}
