variable "project_id" {}

provider "google" {
  credentials = "${file("account.json")}"
  project     = "${var.project_id}"
  region      = "asia-east1"
}

resource "google_storage_bucket" "tmp-bucket" {
  name     = "vod-m3u8-tmp"
  location = "ASIA"
}

resource "google_storage_bucket" "vod-bucket" {
  name     = "vod-m3u8"
  location = "ASIA"
}