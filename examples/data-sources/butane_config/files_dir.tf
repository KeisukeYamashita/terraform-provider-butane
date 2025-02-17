data "butane_config" "config" {
  content   = file("./config.bu")
  files_dir = "./files"
}
