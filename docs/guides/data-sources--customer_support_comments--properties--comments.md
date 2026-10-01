---
page_title: "comments"
subcategory: ""
description: "comments for xcsh_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 2085, "body_sha256": "sha256:7e1e78b1b89cf2b7a8e2ba099498bc52da5108b4b9c9214ec50782e9872f26c5", "canonical_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "child_ids": ["xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info"], "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "parent_id": "xcsh-docs:data-sources:customer_support_comments:reference", "path": "docs/guides/data-sources--customer_support_comments--properties--comments.md", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["comments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/comments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "comments for xcsh_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
- [Property reference](data-sources--customer_support_comments--reference.md)
- comments

<a id="section"></a>

Type: `"list"`. Computed.

List of comments on the customer support ticket.

## Direct properties

<a id="schema-comments--attachment_ids"></a>

### attachment_ids property

Type: `["list", "string"]`. Computed.

Third party ID of any attachment related to this ticket comment.

- [attachments_info](data-sources--customer_support_comments--properties--comments--attachments_info.md): complete subsection reference.

<a id="schema-comments--author_email"></a>

### author_email property

Type: `"string"`. Computed.

Email. Email of the author of the comment.

<a id="schema-comments--author_name"></a>

### author_name property

Type: `"string"`. Computed.

Author. Author of the comment (as a name)

<a id="schema-comments--comment_id"></a>

### comment_id property

Type: `"string"`. Computed.

ID assigned to this comment by support provider.

<a id="schema-comments--created_at"></a>

### created_at property

Type: `"string"`. Computed.

At. Comment creation time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="schema-comments--html"></a>

### html property

Type: `"string"`. Computed.

Comment. Comment body as HTML.

<a id="schema-comments--plain_text"></a>

### plain_text property

Type: `"string"`. Computed.

Comment. Comment body as plain text.

## Next pages

- [comments.attachments_info](data-sources--customer_support_comments--properties--comments--attachments_info.md)
- [Property reference](data-sources--customer_support_comments--reference.md)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
