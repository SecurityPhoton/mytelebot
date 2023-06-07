variable "GOOGLE_PROJECT" {
  type        = string
  description = "GCP project name"
}
           
variable "GOOGLE_REGION" {
  type        = string
  default     = "us-central1-c"
  description = "GCP region to use"
}

variable "GKE_NUM_NODES" {
  type        = number
  description = "Number of nodes in GKE"
  default     = 2
}

variable "GKE_MACHINE_TYPE" {
  type        = string
  default     = "n1-standard-1"
  description = "Machine type"
}