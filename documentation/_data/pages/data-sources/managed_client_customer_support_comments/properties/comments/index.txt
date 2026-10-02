---
page_title: "comments"
subcategory: ""
description: "List of comments on the customer support ticket."
xcsh_docs: {"aliases": ["comments"], "body_bytes": 2513, "body_sha256": "sha256:4604e08320fd7b45719eeb9d9935ea6826c8df7d151de79fd22d0818821e2908", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments:attachments_info"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "parent_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "path": "documentation/data-sources/managed_client_customer_support_comments/properties/comments/index.md", "product": "distributed-cloud", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3331031333232023-3021203022123030-2213020102313132-2331122133113230-0120132330023202-2223010122231213-2330130122103321-3100013203233133", "registry_path": "docs/guides/data-sources--managed_client_customer_support_comments--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["comments"], "schema_version": 1, "sections": [{"aliases": ["attachment ids"], "anchor": "schema-comments--attachment_ids", "description": "Third party ID of any attachment related to this ticket comment.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachment_ids"], "syntax": "attribute", "type": "list"}, {"aliases": ["attachments info"], "anchor": "section", "description": "Information about any attachments (such as screenshots, plain text files) the comment can have.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments:attachments_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["comments", "attachments_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["author email"], "anchor": "schema-comments--author_email", "description": "Email. Email of the author of the comment.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "author_email"], "syntax": "attribute", "type": "string"}, {"aliases": ["author name"], "anchor": "schema-comments--author_name", "description": "Author. Author of the comment (as a name)", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "author_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["comment id"], "anchor": "schema-comments--comment_id", "description": "ID assigned to this comment by support provider.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "comment_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["created at"], "anchor": "schema-comments--created_at", "description": "At. Comment creation time.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "created_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["html"], "anchor": "schema-comments--html", "description": "Comment. Comment body as HTML.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "html"], "syntax": "attribute", "type": "string"}, {"aliases": ["plain text"], "anchor": "schema-comments--plain_text", "description": "Comment. Comment body as plain text.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "plain_text"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/properties/comments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of comments on the customer support ticket.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/)
- comments

<a id="section"></a>

Type: `"list"`. Computed.

List of comments on the customer support ticket.

## Direct properties

<a id="schema-comments--attachment_ids"></a>

### attachment_ids property

Type: `["list", "string"]`. Computed.

Third party ID of any attachment related to this ticket comment.

- [attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/): complete subsection reference.

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

- [comments.attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/)
- [xcsh_managed_client_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/)
