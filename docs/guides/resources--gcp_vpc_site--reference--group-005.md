---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-c9e7c6a6cfa088d705651b09d8aacea1c767f40f4eb5853cdacb39bd1d958778"></a>

## Direct properties — voltstack_cluster.storage_class_list.storage_classes / 8a4645be96e9 / 3

<a id="canonical-47b3bab9615e8f8ec0219bed7f94047b50ff486639d6da99db405638adf669c3"></a>

<a id="canonical-9db1b93f8d6726c2d5f84f43488a0e7f7fdb02ac5390028836e15568aad39df5"></a>

## default_storage_class property — voltstack_cluster.storage_class_list.storage_classes / 8a4645be96e9 / 4

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

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

<a id="canonical-8d677aa2dbbc1bb6ebce673026aec17c437bc869e25f8aa2cae326c8e7a573e4"></a>

<a id="canonical-25a56f09754fc52e9191c1e912f2caecc9534fa11762d0707077385ad1b9660b"></a>

## storage_class_name property — voltstack_cluster.storage_class_list.storage_classes / 8a4645be96e9 / 5

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-21b11ac296fbcd6c788e76c2fe51407250c6e1383e5f8817c2b49f57136a6059"></a>

## Next pages — voltstack_cluster.storage_class_list.storage_classes / 8a4645be96e9 / 6

- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-004.md#canonical-56ffa539204782f5f466a7bc0f01f7b06e0be7bd95c35b7f6529e6d577be27d9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aa42e6ff1d4af638b8301f8d0934d8ee909c62b0c2517535f88e18612eab5fa"></a>

## waf_signatures — waf_signatures / 0ce2ac645b64 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- waf_signatures

<a id="canonical-20cdebd9914ff01019312f079981c5ff72c43afeb869ae02a434b180495dfa66"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-0eee0018bf2a3722a26f355274480b874f8f77d00ad5d01a27945b0a6b4a4742"></a>

## Direct properties — waf_signatures / 0ce2ac645b64 / 3

- [automatic](resources--gcp_vpc_site--reference--group-005.md#canonical-77c3e1df3355dee57cc6ef1eac0c964e50b4c253d833b739d6f44e3b5a894c18): complete subsection reference.

- [manual](resources--gcp_vpc_site--reference--group-005.md#canonical-5a01b16d7961e8470686760d326121c8d159459afd98a6e3bf59651d44a6938a): complete subsection reference.

<a id="canonical-444add2a1fce0de44d6254930ecf36f47b65c2e5f9b46cf6347895fef0fcb96a"></a>

## Next pages — waf_signatures / 0ce2ac645b64 / 4

- [waf_signatures.automatic](resources--gcp_vpc_site--reference--group-005.md#canonical-77c3e1df3355dee57cc6ef1eac0c964e50b4c253d833b739d6f44e3b5a894c18)
- [waf_signatures.manual](resources--gcp_vpc_site--reference--group-005.md#canonical-5a01b16d7961e8470686760d326121c8d159459afd98a6e3bf59651d44a6938a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-77c3e1df3355dee57cc6ef1eac0c964e50b4c253d833b739d6f44e3b5a894c18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dbdd5ad54ac55b61b3b20e58ee8466f27a319222c94b2ceb821cfffb6405e51"></a>

## waf_signatures.automatic — waf_signatures.automatic / 8b37f6fb8176 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0)
- waf_signatures.automatic

<a id="canonical-794d027c6997fac963ad2bf51273c0df9cf5650cdd3ce9388ffaf76f034f8be5"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
automatic = {}
```

<a id="canonical-b8cc2d0d3e9508c1be1e74a820ecc5d8e8f947347162661e9db2cd964f9e3c10"></a>

## Direct properties — waf_signatures.automatic / 8b37f6fb8176 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6b3d99bb5f56c1199a934f656455b39cd761f3dc9e71f631b0ac197f4b78653"></a>

## Next pages — waf_signatures.automatic / 8b37f6fb8176 / 4

- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5a01b16d7961e8470686760d326121c8d159459afd98a6e3bf59651d44a6938a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08460c59233ee156a4aa93186c565f061802c7cd7af1bdeee9c7958f4743fac8"></a>

## waf_signatures.manual — waf_signatures.manual / 37b5c57396b0 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0)
- waf_signatures.manual

<a id="canonical-f1e533a933e3037a95c549327d025ca0fbedfe75c25f86c7bacc27d06c104713"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
manual = {}
```

<a id="canonical-d955e5c75b2f0806b744e26c596309200a2bbe6e2fa297489d46d7feb82bf72e"></a>

## Direct properties — waf_signatures.manual / 37b5c57396b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d57bdea194fbdec09dbf01973714e84a7643e3a73860b6b1dc000cb18a55642"></a>

## Next pages — waf_signatures.manual / 37b5c57396b0 / 4

- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
