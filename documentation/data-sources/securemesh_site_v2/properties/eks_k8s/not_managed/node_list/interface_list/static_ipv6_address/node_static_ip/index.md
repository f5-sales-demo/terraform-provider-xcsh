---
page_title: "eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip"
subcategory: ""
description: "Configure Static IP parameters for a node."
xcsh_docs: {"aliases": ["eks k8s not managed node list interface list static ipv6 address node static ip"], "body_bytes": 4297, "body_sha256": "sha256:fc5d725e55b7dae66eddca3e964d8d6389de9bb6433a690fee2264c8239eb5ad", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ipv6_address", "path": "documentation/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2031112033221333-3321020210013113-0312131122321012-3201233321330322-3100221132002021-3301213022112311-0313233223233210-0033020121301033", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip"], "schema_version": 1, "sections": [{"aliases": ["default gw"], "anchor": "schema-eks_k8s--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--default_gw", "description": "IP address of the default gateway.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip", "default_gw"], "syntax": "attribute", "type": "string"}, {"aliases": ["dns server"], "anchor": "schema-eks_k8s--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--dns_server", "description": "DNS server address for the static interface configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip", "dns_server"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip address"], "anchor": "schema-eks_k8s--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "description": "IP address of the interface and prefix length.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip", "ip_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure Static IP parameters for a node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/)
- [eks_k8s.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/)
- [eks_k8s.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/)
- [eks_k8s.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/static_ipv6_address/)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

## Direct properties

<a id="schema-eks_k8s--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--default_gw"></a>

### default_gw property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-eks_k8s--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--dns_server"></a>

### dns_server property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

## Next pages

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/static_ipv6_address/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
