---
page_title: "xcsh_public_ip_binding Resource"
description: |-
  Manage regional virtual-site availability of an existing allocated public IP.
---

# xcsh_public_ip_binding (Resource)

Adopts only the binding of an existing public IP object. It changes the regional
virtual-site reference and restores the captured original binding on delete.
It never allocates, deallocates or deletes the public IP.

The target virtual site must be `REGIONAL_EDGE`. Reference this binding's name
and namespace from an HTTP load balancer's
`advertise_custom.advertise_where.advertise_on_public.public_ip` block.

Keep one owner for each public IP binding. An unexpected IP address fails.
Refresh observes changed bindings so Terraform can repair drift. Delete refuses
to overwrite a binding changed by another owner. An uncertain PUT is reconciled
by exact GET before another mutation.

## Example Usage

```terraform
resource "xcsh_public_ip_binding" "canada" {
  name                   = var.allocated_public_ip_name
  namespace              = "shared"
  expected_ip            = var.allocated_public_ip
  virtual_site           = xcsh_virtual_site.canada.name
  virtual_site_namespace = xcsh_virtual_site.canada.namespace
}
```

## Schema

### Required

- `name` (String) Existing public IP object name; changes replace the binding.
- `namespace` (String) Existing public IP namespace; changes replace the binding.
- `expected_ip` (String) Exact allocated IPv4/IPv6 address; changes replace the binding.
- `virtual_site` (String) Regional Edge virtual site.
- `virtual_site_namespace` (String) Namespace of that virtual site.

### Read-Only

- `id` (String) Namespace/name identity.
- `original_bindings` (String) Original binding JSON restored on delete.
- `managed_bindings` (String) Last applied binding JSON used for ownership checks.
