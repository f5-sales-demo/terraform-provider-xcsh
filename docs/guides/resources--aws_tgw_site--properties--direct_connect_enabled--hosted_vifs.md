---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: ""
description: "direct_connect_enabled.hosted_vifs for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2224, "body_sha256": "sha256:b8925f35b878de2305befc2f5a8ea395b6c97ccdf67e8ccfa6ea5c074c0f2a8b", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "path": "docs/guides/resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled.hosted_vifs for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# direct_connect_enabled.hosted_vifs

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md)
- direct_connect_enabled.hosted_vifs

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_registration_over_direct_connect",
    "site_registration_over_internet")}
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
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

Terraform syntax:

```terraform
hosted_vifs {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site_registration_over_direct_connect](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md): complete subsection reference.

- [site_registration_over_internet](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md): complete subsection reference.

- [vif_list](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md): complete subsection reference.

## Next pages

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md)
- [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
