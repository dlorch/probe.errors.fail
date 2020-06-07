terraform {
  backend "gcs" {
    bucket = "errors-fail-terraform-state"
    prefix = "probe.errors.fail"
  }
}
