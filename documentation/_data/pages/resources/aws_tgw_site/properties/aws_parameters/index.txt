---
page_title: "aws_parameters"
subcategory: ""
description: "aws_parameters for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 15024, "body_sha256": "sha256:c9ebf7ab86438aa2eb141b89ff1d46259449ac7ac19dd8231abf12e1e1bb90be", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:aws_cred", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:custom_security_group", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_encryption", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_internet_vip", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_internet_vip", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:existing_tgw", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:f5xc_security_group", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:reserved_tgw_cidr", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:tgw_cidr"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["aws_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- aws_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Setup AWS services VPC, transit gateway and site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "az_nodes",
    "instance_type",
    "ssh_key"),
  validators.ConflictingObjectAttributes("custom_security_group",
    "f5xc_security_group"),
  validators.ConflictingObjectAttributes("disable_encryption",
    "enable_encryption"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip"),
  validators.ConflictingObjectAttributes("existing_tgw",
    "new_tgw"),
  validators.ConflictingObjectAttributes("new_vpc",
    "vpc_id"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "nodes_per_az"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "total_nodes"),
  validators.ConflictingObjectAttributes("nodes_per_az",
    "total_nodes"),
  validators.ConflictingObjectAttributes("reserved_tgw_cidr",
    "tgw_cidr")}
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
  "x-ves-oneof-field-deployment": "[\"aws_cred\"]",
  "x-ves-oneof-field-encryption_choice": "[\"disable_encryption\",\"enable_encryption\"]",
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]",
  "x-ves-oneof-field-security_group_choice": "[\"custom_security_group\",\"f5xc_security_group\"]",
  "x-ves-oneof-field-service_vpc_choice": "[\"new_vpc\",\"vpc_id\"]",
  "x-ves-oneof-field-tgw_choice": "[\"existing_tgw\",\"new_tgw\"]",
  "x-ves-oneof-field-tgw_cidr_choice": "[\"reserved_tgw_cidr\",\"tgw_cidr\"]",
  "x-ves-oneof-field-worker_nodes": "[\"no_worker_nodes\",\"nodes_per_az\",\"total_nodes\"]"
}
```

Terraform syntax:

```terraform
aws_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/): complete subsection reference.

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/): complete subsection reference.

<a id="schema-aws_parameters--aws_region"></a>

### aws_region property

Type: `"string"`. Optional.

AWS Region of your services VPC, where F5XC site will be deployed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/): complete subsection reference.

- [custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/custom_security_group/): complete subsection reference.

- [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_encryption/): complete subsection reference.

- [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_internet_vip/): complete subsection reference.

<a id="schema-aws_parameters--disk_size"></a>

### disk_size property

Type: `"number"`. Optional.

Node disk size for all node in the F5XC site. Unit is GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/): complete subsection reference.

- [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_internet_vip/): complete subsection reference.

- [existing_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/): complete subsection reference.

- [f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/f5xc_security_group/): complete subsection reference.

<a id="schema-aws_parameters--instance_type"></a>

### instance_type property

Type: `"string"`. Optional.

AWS Instance Type for Node. Instance size based on the performance.

Upstream description:

Instance size based on the performance.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/): complete subsection reference.

- [new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/): complete subsection reference.

- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/no_worker_nodes/): complete subsection reference.

<a id="schema-aws_parameters--nodes_per_az"></a>

### nodes_per_az property

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [reserved_tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/reserved_tgw_cidr/): complete subsection reference.

<a id="schema-aws_parameters--ssh_key"></a>

### ssh_key property

Type: `"string"`. Optional.

Public SSH key for accessing nodes of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/tgw_cidr/): complete subsection reference.

<a id="schema-aws_parameters--total_nodes"></a>

### total_nodes property

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

<a id="schema-aws_parameters--vpc_id"></a>

### vpc_id property

Type: `"string"`. Optional.

Exclusive with \[new\_vpc\] Existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Existing VPC ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [aws_parameters.admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/)
- [aws_parameters.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/)
- [aws_parameters.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/)
- [aws_parameters.custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/custom_security_group/)
- [aws_parameters.disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_encryption/)
- [aws_parameters.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_internet_vip/)
- [aws_parameters.enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/)
- [aws_parameters.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_internet_vip/)
- [aws_parameters.existing_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/)
- [aws_parameters.f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/f5xc_security_group/)
- [aws_parameters.new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/)
- [aws_parameters.new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/)
- [aws_parameters.no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/no_worker_nodes/)
- [aws_parameters.reserved_tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/reserved_tgw_cidr/)
- [aws_parameters.tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/tgw_cidr/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
