---
page_title: "filters.global_filters"
subcategory: ""
description: "Global Filters. List of global filters."
xcsh_docs: {"aliases": ["filters global filters"], "body_bytes": 5424, "body_sha256": "sha256:6d22d3dbda23570842f850202163247a4b8ab9c8530a8a8d399ec0c7549db9ce", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:properties:filters:global_filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:filters", "path": "documentation/data-sources/device_intelligence_devices/properties/filters/global_filters/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2213122233032232-1123233111303120-1223001233111200-2230310213022301-2321112010220310-2230201122130203-3110022222321122-3113121101123113", "registry_path": "docs/guides/data-sources--device_intelligence_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filters", "global_filters"], "schema_version": 1, "sections": [{"aliases": ["filters global filters key"], "anchor": "schema-filters--global_filters--key", "description": "Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible values are `TIMESTAMP`, `USERNAME`, `CLIENT_TOKEN`, `IP_ADDRESS`, `ASN`, `AS_ORGANIZATION`, `COUNTRY`, `METHOD`, `HOST`, `PATH`, `URL`, `REFERER`, `TRAFFIC_CHANNEL`, `IS_ATTACK`, `BOT_REASON`, `TRAFFIC_TYPE`, `THREAT_TYPE`,", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:filters:global_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "global_filters", "key"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters global filters op"], "anchor": "schema-filters--global_filters--op", "description": "Operator for query filter - IN: Filter Operator Specifies that query result includes filter values - NOT_IN: Filter Operator Specifies that query result excludes filter values - MATCHES_REGEX: Filter Operator Specifies that query result matches filter regex - DOES_NOT_MATCH_REGEX: Filter.. Possible values are `IN`,", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:filters:global_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "global_filters", "op"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters global filters values"], "anchor": "schema-filters--global_filters--values", "description": "Values. An unordered list of filter strings.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:filters:global_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "global_filters", "values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/filters/global_filters/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Global Filters. List of global filters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filters.global_filters

Breadcrumbs:

- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/)
- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/)
- filters.global_filters

<a id="section"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

## Direct properties

<a id="schema-filters--global_filters--key"></a>

### key property

Type: `"string"`. Optional.

\[Enum:
TIMESTAMP|USERNAME|CLIENT\_TOKEN|IP\_ADDRESS|ASN|AS\_ORGANIZATION|COUNTRY|METHOD|HOST|PATH|URL|REFERER|TRAFFIC\_CHANNEL|IS\_ATTACK|BOT\_REASON|TRAFFIC\_TYPE|THREAT\_TYPE|SDK\_VERSION|ACTION\_TAKEN|COOKIE\_AGE|BOT\_COOKIE|USER\_AGENT|USER\_AGENT\_OS\_FAMILY|USER\_AGENT\_FAMILY|BROWSER\_FINGERPRINT|USER\_FINGERPRINT|HEADER\_FINGERPRINT|DEVICE\_ID|FLOW|AGENT|APPLICATION\_NAME|PROTECTED\_APPLICATION|RESPONSE\_CODE|SERVER\_RESPONSE\_CODE|TRANSACTION\_RESULT|MOBILE\_TRANSACTION\_INSIGHT|WEB\_TRANSACTION\_INSIGHT|TRIGGERED\_RULE|FLOW\_CATEGORY|FLOW\_LABEL|ENDPOINT\_NAME|ENDPOINT\_LABEL|BOT\_ENDPOINT\_POLICY|KNOWN\_BOT\_NAME|KNOWN\_BOT\_CATEGORY|KNOWN\_BOT\_PROVIDER|KNOWN\_BOT\_CATEGORY\_TYPE|KNOWN\_BOT\_MITIGATION|ABSOLUTE|PERCENTAGE|TREND|ENDPOINT\_POLICY\]
Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible
values are \`TIMESTAMP\`, \`USERNAME\`, \`CLIENT\_TOKEN\`, \`IP\_ADDRESS\`, \`ASN\`,
\`AS\_ORGANIZATION\`, \`COUNTRY\`, \`METHOD\`, \`HOST\`, \`PATH\`, \`URL\`, \`REFERER\`,
\`TRAFFIC\_CHANNEL\`, \`IS\_ATTACK\`, \`BOT\_REASON\`, \`TRAFFIC\_TYPE\`, \`THREAT\_TYPE\`,
\`SDK\_VERSION\`, \`ACTION\_TAKEN\`, \`COOKIE\_AGE\`, \`BOT\_COOKIE\`, \`USER\_AGENT\`,
\`USER\_AGENT\_OS\_FAMILY\`, \`USER\_AGENT\_FAMILY\`, \`BROWSER\_FINGERPRINT\`,
\`USER\_FINGERPRINT\`, \`HEADER\_FINGERPRINT\`, \`DEVICE\_ID\`, \`FLOW\`, \`AGENT\`,
\`APPLICATION\_NAME\`, \`PROTECTED\_APPLICATION\`, \`RESPONSE\_CODE\`, \`SERVER\_RESPONSE\_CODE\`,
\`TRANSACTION\_RESULT\`, \`MOBILE\_TRANSACTION\_INSIGHT\`, \`WEB\_TRANSACTION\_INSIGHT\`,
\`TRIGGERED\_RULE\`, \`FLOW\_CATEGORY\`, \`FLOW\_LABEL\`, \`ENDPOINT\_NAME\`, \`ENDPOINT\_LABEL\`,
\`BOT\_ENDPOINT\_POLICY\`, \`KNOWN\_BOT\_NAME\`, \`KNOWN\_BOT\_CATEGORY\`, \`KNOWN\_BOT\_PROVIDER\`,
\`KNOWN\_BOT\_CATEGORY\_TYPE\`, \`KNOWN\_BOT\_MITIGATION\`, \`ABSOLUTE\`, \`PERCENTAGE\`, \`TREND\`,
\`ENDPOINT\_POLICY\`. Defaults to \`TIMESTAMP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TIMESTAMP",
    "USERNAME",
    "CLIENT_TOKEN",
    "IP_ADDRESS",
    "ASN",
    "AS_ORGANIZATION",
    "COUNTRY",
    "METHOD",
    "HOST",
    "PATH",
    "URL",
    "REFERER",
    "TRAFFIC_CHANNEL",
    "IS_ATTACK",
    "BOT_REASON",
    "TRAFFIC_TYPE",
    "THREAT_TYPE",
    "SDK_VERSION",
    "ACTION_TAKEN",
    "COOKIE_AGE",
    "BOT_COOKIE",
    "USER_AGENT",
    "USER_AGENT_OS_FAMILY",
    "USER_AGENT_FAMILY",
    "BROWSER_FINGERPRINT",
    "USER_FINGERPRINT",
    "HEADER_FINGERPRINT",
    "DEVICE_ID",
    "FLOW",
    "AGENT",
    "APPLICATION_NAME",
    "PROTECTED_APPLICATION",
    "RESPONSE_CODE",
    "SERVER_RESPONSE_CODE",
    "TRANSACTION_RESULT",
    "MOBILE_TRANSACTION_INSIGHT",
    "WEB_TRANSACTION_INSIGHT",
    "TRIGGERED_RULE",
    "FLOW_CATEGORY",
    "FLOW_LABEL",
    "ENDPOINT_NAME",
    "ENDPOINT_LABEL",
    "BOT_ENDPOINT_POLICY",
    "KNOWN_BOT_NAME",
    "KNOWN_BOT_CATEGORY",
    "KNOWN_BOT_PROVIDER",
    "KNOWN_BOT_CATEGORY_TYPE",
    "KNOWN_BOT_MITIGATION",
    "ABSOLUTE",
    "PERCENTAGE",
    "TREND",
    "ENDPOINT_POLICY"),
}
```

<a id="schema-filters--global_filters--op"></a>

### op property

Type: `"string"`. Optional.

\[Enum:
IN|NOT\_IN|MATCHES\_REGEX|DOES\_NOT\_MATCH\_REGEX|INCLUDES|DOES\_NOT\_INCLUDE|STARTS\_WITH|ENDS\_WITH\]
Operator for query filter - IN: Filter Operator Specifies that query result includes filter values -
NOT\_IN: Filter Operator Specifies that query result excludes filter values - MATCHES\_REGEX: Filter
Operator Specifies that query result matches filter regex - DOES\_NOT\_MATCH\_REGEX: Filter..
Possible values are \`IN\`, \`NOT\_IN\`, \`MATCHES\_REGEX\`, \`DOES\_NOT\_MATCH\_REGEX\`,
\`INCLUDES\`, \`DOES\_NOT\_INCLUDE\`, \`STARTS\_WITH\`, \`ENDS\_WITH\`. Defaults to \`IN\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IN",
    "NOT_IN",
    "MATCHES_REGEX",
    "DOES_NOT_MATCH_REGEX",
    "INCLUDES",
    "DOES_NOT_INCLUDE",
    "STARTS_WITH",
    "ENDS_WITH"),
}
```

<a id="schema-filters--global_filters--values"></a>

### values property

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

## Next pages

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/)
- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
