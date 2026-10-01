---
page_title: "dns_ntp_config"
subcategory: ""
description: "dns_ntp_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2843, "body_sha256": "sha256:81eccd4b224490192711ef32d0744dd01d7eae2f5b890f258ac1d521d4181aee", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/dns_ntp_config/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["dns_ntp_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/dns_ntp_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dns_ntp_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- dns_ntp_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_dns",
    "f5_dns_default"),
  validators.ConflictingObjectAttributes("custom_ntp",
    "f5_ntp_default")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_server_choice": "[\"custom_dns\",\"f5_dns_default\"]",
  "x-ves-oneof-field-ntp_server_choice": "[\"custom_ntp\",\"f5_ntp_default\"]"
}
```

Terraform syntax:

```terraform
dns_ntp_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/): complete subsection reference.

- [custom_ntp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/): complete subsection reference.

- [f5_dns_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/f5_dns_default/): complete subsection reference.

- [f5_ntp_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/f5_ntp_default/): complete subsection reference.

## Next pages

- [dns_ntp_config.custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/)
- [dns_ntp_config.custom_ntp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/)
- [dns_ntp_config.f5_dns_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/f5_dns_default/)
- [dns_ntp_config.f5_ntp_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/f5_ntp_default/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
