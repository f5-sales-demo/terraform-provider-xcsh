resource "xcsh_public_ip_binding" "canada" {
  name                   = var.allocated_public_ip_name
  namespace              = "shared"
  expected_ip            = var.allocated_public_ip
  virtual_site           = xcsh_virtual_site.canada.name
  virtual_site_namespace = xcsh_virtual_site.canada.namespace
}
