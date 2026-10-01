---
page_title: "custom_security_group"
subcategory: "Infrastructure"
description: "custom_security_group for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3840, "body_sha256": "sha256:b563113d1f1759a7ad58527752ea634fb0eafbf9d2f3e196737ceeed792b3967", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:custom_security_group", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:custom_security_group", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--custom_security_group.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_security_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/custom_security_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_security_group for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_security_group

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- custom_security_group

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_security\_group, f5xc\_security\_group\] Enter pre created security groups for
slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on
existing VPC.

Upstream description:

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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

OneOf alternatives in this subsection:

- [custom_security_group](resources--aws_vpc_site--properties--custom_security_group.md#section)
- [f5xc_security_group](resources--aws_vpc_site--properties--f5xc_security_group.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_security_group {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_security_group--inside_security_group_id"></a>

### inside_security_group_id property

Type: `"string"`. Optional.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="schema-custom_security_group--outside_security_group_id"></a>

### outside_security_group_id property

Type: `"string"`. Optional.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

## Next pages

- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
