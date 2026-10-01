---
page_title: "aws.byoc.connections"
subcategory: ""
description: "aws.byoc.connections for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 12401, "body_sha256": "sha256:ab958b904d77ee4af790ae8a9b152fc4cb3917d68a33542a79e21781c137647f", "canonical_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:auth_key", "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:ipv4", "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:metadata", "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:system_generated_name"], "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections", "parent_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc", "path": "docs/guides/data-sources--cloud_link--properties--aws--byoc--connections.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws", "byoc", "connections"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/aws/byoc/connections/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws.byoc.connections for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md)
- [Property reference](data-sources--cloud_link--reference.md)
- [aws](data-sources--cloud_link--properties--aws.md)
- [aws.byoc](data-sources--cloud_link--properties--aws--byoc.md)
- aws.byoc.connections

<a id="section"></a>

Type: `"list"`. Computed.

List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but
will be used for connecting sites and REs.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [auth_key](data-sources--cloud_link--properties--aws--byoc--connections--auth_key.md): complete subsection reference.

<a id="schema-aws--byoc--connections--bgp_asn"></a>

### bgp_asn property

Type: `"number"`. Computed.

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Upstream description:

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-aws--byoc--connections--connection_id"></a>

### connection_id property

Type: `"string"`. Computed.

ID of the existing AWS Direct Connect Connection.

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
    "pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [ipv4](data-sources--cloud_link--properties--aws--byoc--connections--ipv4.md): complete subsection reference.

- [metadata](data-sources--cloud_link--properties--aws--byoc--connections--metadata.md): complete subsection reference.

<a id="schema-aws--byoc--connections--region"></a>

### region property

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
Region. Region where the connection is setup. Possible values are \`ap-northeast-1\`,
\`ap-southeast-1\`, \`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`,
\`us-east-2\`, \`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`,
\`ap-northeast-2\`, \`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`,
\`me-south-1\`, \`us-west-1\`, \`ap-southeast-3\`.

Upstream description:

Region where the connection is setup.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [system_generated_name](data-sources--cloud_link--properties--aws--byoc--connections--system_generated_name.md): complete subsection reference.

<a id="schema-aws--byoc--connections--tags"></a>

### tags property

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-aws--byoc--connections--user_assigned_name"></a>

### user_assigned_name property

Type: `"string"`. Computed.

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Upstream description:

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-aws--byoc--connections--virtual_interface_type"></a>

### virtual_interface_type property

Type: `"string"`. Computed.

\[Enum: PRIVATE\] Defines the type of virtual interface that needs to be configured on AWS -
PRIVATE: Private A private virtual interface should be used to access an Amazon VPC using private IP
addresses. - TRANSIT: Transit A transit virtual interface is a VLAN that transports traffic from a
Direct Connect.. The only possible value is \`PRIVATE\`. Defaults to \`PRIVATE\`.

Upstream description:

Defines the type of virtual interface that needs to be configured on AWS

&#8203;- PRIVATE: Private

A private virtual interface should be used to access an Amazon VPC using private IP addresses.
&#8203;- TRANSIT: Transit

A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one
or more transit gateways.

Receipt-pinned upstream constraints:

```json
{
  "default": "PRIVATE",
  "enum": [
    "PRIVATE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-aws--byoc--connections--vlan"></a>

### vlan property

Type: `"number"`. Computed.

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Upstream description:

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4094,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  }
}
```

## Next pages

- [aws.byoc.connections.auth_key](data-sources--cloud_link--properties--aws--byoc--connections--auth_key.md)
- [aws.byoc.connections.ipv4](data-sources--cloud_link--properties--aws--byoc--connections--ipv4.md)
- [aws.byoc.connections.metadata](data-sources--cloud_link--properties--aws--byoc--connections--metadata.md)
- [aws.byoc.connections.system_generated_name](data-sources--cloud_link--properties--aws--byoc--connections--system_generated_name.md)
- [aws.byoc](data-sources--cloud_link--properties--aws--byoc.md)
- [xcsh_cloud_link](../data-sources/cloud_link.md)
