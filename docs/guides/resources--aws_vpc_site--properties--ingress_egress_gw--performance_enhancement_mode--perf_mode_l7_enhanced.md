---
page_title: "ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: "Infrastructure"
description: "ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2338, "body_sha256": "sha256:83f9b76b0bb3feeb14db32706a07bca9435d22df225c3a3ada7195474f04c940", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode.md)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

## Direct properties

- [jumbo_disabled](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md): complete subsection reference.

- [jumbo_enabled](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
