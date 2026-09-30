---
page_title: "performance_enhancement_mode.perf_mode_l3_enhanced"
subcategory: ""
description: "performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1852, "body_sha256": "sha256:094f95360a3e564af7dac992b536621699585374c714f6c389796353c095b77b", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode", "path": "docs/guides/resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l3_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# performance_enhancement_mode.perf_mode_l3_enhanced

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [performance_enhancement_mode](resources--aws_tgw_site--properties--performance_enhancement_mode.md)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

## Direct properties

- [jumbo](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md): complete subsection reference.

- [no_jumbo](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md)
- [performance_enhancement_mode](resources--aws_tgw_site--properties--performance_enhancement_mode.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
