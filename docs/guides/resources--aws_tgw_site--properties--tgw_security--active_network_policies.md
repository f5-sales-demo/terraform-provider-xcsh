---
page_title: "tgw_security.active_network_policies"
subcategory: ""
description: "tgw_security.active_network_policies for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1466, "body_sha256": "sha256:d886072411da6ae010bfccf998f10c189433fa52ea4ccef445e22c1ee2871259", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "path": "docs/guides/resources--aws_tgw_site--properties--tgw_security--active_network_policies.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security", "active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security.active_network_policies for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.active_network_policies

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- tgw_security.active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](resources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [tgw_security.active_network_policies.network_policies](resources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md)
- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
