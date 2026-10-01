---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-b66f9e99fcde55cdb33daa16d3fcf57f1d1a2f73e92351f95b82fdefbc9d3fc5"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public / edcc7e053bfc / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcf8d7e41745566c9833f921810f02c1749981f95e69aa00aa9bd1d4393184e4"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports — stateful_service.advertise_options.advertise_on_public.multi_ports / 9b7f1b880656 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- stateful_service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-9bdcbad2e908f235af1c6e60ce3729852ce28a1b5939970fda519cea8b932819"></a>

Type: `"object"`. single nested block, Optional.

Advertise Multiple Ports. Advertise multiple ports.

Upstream description:

Advertise multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
```

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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2ba114b675425a2481ae72f8993ae7c076d6c6460bd6471b5b46630a2aaa6e8c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports / 9b7f1b880656 / 3

- [ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f): complete subsection reference.

<a id="canonical-e41d67e50843936ed70a8a489002bbae0327e1e2008191e4821c8b059c021bdf"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports / 9b7f1b880656 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf8e8e357225d4ebb17c05f93c72d99d5fa472a5129f36feeeef6a33986ef2b1"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports — stateful_service.advertise_options.advertise_on_public.multi_ports.ports / b62f1eeb83dd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-dac7c5f565f184b839af043702d21223b657ca3961be32417e9fb578bf292d04"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c8affd089fdc34139baed12a795d0b707ded0f2031771b7fc97635d6fb09ecd"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports / b62f1eeb83dd / 3

- [http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687): complete subsection reference.

- [port](resources--workload--reference--group-024.md#canonical-4f57cb95ce001efb49dfd84bf08143f0bdbdf929271f32aff825a054eafdfa5d): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-024.md#canonical-3561ccb026c2d7060c92d2f7b922913f055e6878453c25bf5c512cf37a84e927): complete subsection reference.

<a id="canonical-589b82094e921d9e8bc935a750d6ef7a2a70e159844ebc08d8965f9f21b50a8c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports / b62f1eeb83dd / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-4f57cb95ce001efb49dfd84bf08143f0bdbdf929271f32aff825a054eafdfa5d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer](resources--workload--reference--group-024.md#canonical-3561ccb026c2d7060c92d2f7b922913f055e6878453c25bf5c512cf37a84e927)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba4ca210983e99e603b304e16f7c4e9d45fc8b5875b7a96acbaaf3c56200353e"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b0687718c055 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-b1d4854a016581e0dfb5825d32ea445477983f5d13f5e33c91097f9511e9e2f2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-051e586958dcf44cbd685c35a23a41f896dd60446cf96e7ae93e7f80cbfd206d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b0687718c055 / 3

- [default_route](resources--workload--reference--group-021.md#canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9): complete subsection reference.

<a id="canonical-512b56c4917c1b8ebe064dec1d7eeb813d2c679a82c3fdc2e4694cab65a0181a"></a>

<a id="canonical-500a55e4e38d7c5464dc1d6ee387e9c85fd10b9938291a9f3ee6970334ab61b9"></a>

## domains property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b0687718c055 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](resources--workload--reference--group-021.md#canonical-56c092e0c3b227f75a59bce9b7b16f4288a3e04eda7b846d0b69486fa1dc7c91): complete subsection reference.

- [https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40): complete subsection reference.

- [specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33): complete subsection reference.

<a id="canonical-ef44541a1dd36608f134a5242dd2dcd5c84fc356af7328380ef7d6c3c8264229"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b0687718c055 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](resources--workload--reference--group-021.md#canonical-56c092e0c3b227f75a59bce9b7b16f4288a3e04eda7b846d0b69486fa1dc7c91)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d11dbe8e0e85cb91c4afa54f023955ca34c13ebe09e563c75850686af17e609"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 02fe855dfd87 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-66105aa5a64ce25de1141226c473b1e3555aa0130a44bb98462f2a22e15ebd51"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c12c7cda665663141fb53c40cb5179e7645017e39b4f8748be8e55541f40218"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 02fe855dfd87 / 3

- [auto_host_rewrite](resources--workload--reference--group-021.md#canonical-1ddc925c433efecac3c2b355e821284b3c8f0be2b0eb68ac13c88210d921f29d): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-021.md#canonical-78f56901322c12ef9a8efcab733219ae4a539b84e99fc1dd8f9e2c075334c5cd): complete subsection reference.

<a id="canonical-07a87a6025323859ad80b2d25af3ed30d6f6cc6d0d23fcef4e675cbda8245843"></a>

<a id="canonical-709917c49aaf07d513543db4af39f3db96869c6958ee96fc84c8370437d9432e"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 02fe855dfd87 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-43f6d406a4d7f605c201902b468fc88df046b5c11eb8fc6ad85e52b81b70019e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 02fe855dfd87 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-021.md#canonical-1ddc925c433efecac3c2b355e821284b3c8f0be2b0eb68ac13c88210d921f29d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-021.md#canonical-78f56901322c12ef9a8efcab733219ae4a539b84e99fc1dd8f9e2c075334c5cd)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1ddc925c433efecac3c2b355e821284b3c8f0be2b0eb68ac13c88210d921f29d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c35051816224f3eb39a34e4b0f5553a27cac3db341d972c6eae582faffbac2f8"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / efb195eb344d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-f1928be105d3fea07d1b86e2e8c1c599e9a3e0de164a961371d14437a2537a4f"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-3d3298628cd2e31da1922b55921c61b31796dc11f7970abc4707fb1bd7e4601b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / efb195eb344d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b22fec83a90c44baa620c4ca77a132f6fb67d871495eddd05a4bf3db4aba8ae"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / efb195eb344d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-78f56901322c12ef9a8efcab733219ae4a539b84e99fc1dd8f9e2c075334c5cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-100b2c4090a85d9a3fcbefd081472f1297984acb90531734ac52b442c75cbe49"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e1c13377d4d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-e6e95db8855d5e67e97562062b5463371f397ac62220f240a7578807e29f929e"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-3c2488138e55a7f9e66283350e8e5821658888391253495f53a1a1336bb11a0d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e1c13377d4d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b46eab1fbe647c3410ad0e2aad6e0367e84030c836c5f8190354957de28a637d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e1c13377d4d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-93660dd6a6f7ad58e6a2737080bdcd2e749051ace6e80df5f269c44fc491c8d9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-56c092e0c3b227f75a59bce9b7b16f4288a3e04eda7b846d0b69486fa1dc7c91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbb6075048568356c1f029f4c906519c2fa2cef86217ba80dfc6c6551f18d1b1"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1dc275a7f2f3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-83718bbfe71564cbb0ca62079be0534cb2e3a18da8492cc8b3dbf548bb7e674f"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-68daa66bb2e509a72e223566f2d5b7c4273fdb007aeafca17145b80d1473721a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1dc275a7f2f3 / 3

<a id="canonical-d7c33a93c23ec21867820e8a3e2a80570a92513195081fc90981d2da5513910a"></a>

<a id="canonical-7a4c72c012e235328ac1674376a92043d86959273e4675b990b360c6f7c2c8d7"></a>

## dns_volterra_managed property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1dc275a7f2f3 / 4

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-232e8c89a54861e009d3e37bfbb9c74f897b1ecc4173cb21ef3ca8d0ad7000d9"></a>

<a id="canonical-cb80e5bca729fb3fb70b8576e6cdc73a03cd70735a41fe97c00e37104379e3b6"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1dc275a7f2f3 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-cb0736e833b4dcd8376f788a764dbf23417bd09ad0b1b297282f14f9ecd111d9"></a>

<a id="canonical-0cf85d5373af2e158d653d74352608c15c056d53d36c46415644997cc11619d3"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1dc275a7f2f3 / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-419a45896d7c49c1327b96304d759a015cefe4ddcf37a6c929d4826ce1ce982d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1dc275a7f2f3 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68841884ec4458eee7d4be182dd06b21bb56d207e7d95fb365ba13fa6b0b711f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-53df8f89fcbae0d363a6fb2a24795781e3b4fa1d15e65976682b07dab11beb3a"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-49b027a1c6f49fdc2267500add055903fe9e5411c886827c490d752eb0a368d9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 3

<a id="canonical-4ca36d5a73a7b5c7fadcab50eebdb5cbca13b801ed372e19e26e8e1cf3619b99"></a>

<a id="canonical-b2da6053bbcd5ba5151d633d2f7092ead1b9120e361f680e44f5ce6b1cf192c8"></a>

## add_hsts property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 4

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-d3fc45a3f0609032d748c5b5e5abd632cec529af0bc363e957512737852f96a5"></a>

<a id="canonical-9335f66593b95901f75062b6b27742c7b002a2a9b5d166205d824248b5b24fd3"></a>

## append_server_name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-021.md#canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8): complete subsection reference.

<a id="canonical-0458b802a5f76f69b3c1159b2b178c5ecf78d27b8f56927292ea33041f4dcd09"></a>

<a id="canonical-957e9a3fc37c47e2cb4071ea04757ddab460f0fbb51131f65953c7c1c081bbd3"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-021.md#canonical-f2946c4e86f3c1741712e882a1e60d976932bcfcb3e7f415690ea6a3edddc9ea): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-021.md#canonical-45936111480a3a4be47fa0c5096d7c0a50f14cf1d47d0c272fc14753a6e753bd): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-021.md#canonical-907f50f864e14ec5a38d72f62a0e863830ba537307790091d05fee3d4baef8ba): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-021.md#canonical-07a49a97b563f697aaf5914f6252ede0564f603ba2edcdd289c645d2b380d260): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50): complete subsection reference.

<a id="canonical-d39079323d32b1d93963ad286d21e81c486c2cf58683b6f5dd356d26682f7813"></a>

<a id="canonical-86d0f3028f1e662f4128a778f6d591ec5640aa310af3f20979d6f7abcd745a7b"></a>

## http_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 7

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](resources--workload--reference--group-021.md#canonical-65d5ef0a35d63014f59cc3c1f97bdd4784c7ff7a0fa9a69c031410d774bb9923): complete subsection reference.

- [pass_through](resources--workload--reference--group-021.md#canonical-b31dda061c4a7caa33f3654faca7ef2c566e9159cda457e0250aef367bbb2dca): complete subsection reference.

<a id="canonical-f60b1dbc338bc0dfbc12fc9423cfbba4eb1aeb5462392bf8d24d7915253c97e9"></a>

<a id="canonical-c2d9bc7e971716224c8b5b6936931725594953fabcafe9ac95c51a446c531011"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-a687ea8b2dc88be07ddf9885132d23e5369a5a37ac4c93428bba16fb2b773c88"></a>

<a id="canonical-ab6b586f886b49cae0e47fd6a3f8a480b145016377ae9e0294cc9833e9fa17b7"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-c36a388ae7b5f6dc7d3736bd4fe60b617ff810a33e5fb234b2cad42abd9f7a89"></a>

<a id="canonical-d8e07e4e5dfd9d03e0824e20938025f227e8fce0fbbc082275862785579478aa"></a>

## server_name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5): complete subsection reference.

<a id="canonical-008b9403e9388463230d7904a877b97e689c14e141e9e238598e2d11e6d53c48"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 306a55c1d638 / 11

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header](resources--workload--reference--group-021.md#canonical-f2946c4e86f3c1741712e882a1e60d976932bcfcb3e7f415690ea6a3edddc9ea)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-021.md#canonical-45936111480a3a4be47fa0c5096d7c0a50f14cf1d47d0c272fc14753a6e753bd)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-021.md#canonical-907f50f864e14ec5a38d72f62a0e863830ba537307790091d05fee3d4baef8ba)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-021.md#canonical-07a49a97b563f697aaf5914f6252ede0564f603ba2edcdd289c645d2b380d260)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-021.md#canonical-65d5ef0a35d63014f59cc3c1f97bdd4784c7ff7a0fa9a69c031410d774bb9923)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through](resources--workload--reference--group-021.md#canonical-b31dda061c4a7caa33f3654faca7ef2c566e9159cda457e0250aef367bbb2dca)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53a3a94fddf016c193860f545af964195158c8dc3adbadf7e3e2db5de106e4f0"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 992a816e4424 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-8b72a91faaeb27cd95f770b184c143c095bd4aa6913602600c6d3d4465a10d47"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7ee305bb864c148d8f4ef2eddcb9421b33254c1bfd04bbba5b188832e9caad6"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 992a816e4424 / 3

- [default_coalescing](resources--workload--reference--group-021.md#canonical-3b1ec9cd5e4a0dd07342b6d982895b40a559a9b478c97a3bdc4a6c454d4215c0): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-021.md#canonical-bf08e330d12c0366bdeabcb9d81fd6d82b608e25b5239968c77d81f4ecd24979): complete subsection reference.

<a id="canonical-ea1fda15a6925185c9725d661301da6ef80af49e76725e8d9e6368caeafa41b1"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 992a816e4424 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-021.md#canonical-3b1ec9cd5e4a0dd07342b6d982895b40a559a9b478c97a3bdc4a6c454d4215c0)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-021.md#canonical-bf08e330d12c0366bdeabcb9d81fd6d82b608e25b5239968c77d81f4ecd24979)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3b1ec9cd5e4a0dd07342b6d982895b40a559a9b478c97a3bdc4a6c454d4215c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1894614c796cec12ce507f13229e134633662c6818f86a9853d9e3a7705f846a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 00ed7b4dd543 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-4553b231a58db641e50049864c804ee8b80836343fc66082307e90ba32b4a0b4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

<a id="canonical-155b45aa62e4f534d90980898e0e249a12da7b9343626129b5a5a973ffbaa237"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 00ed7b4dd543 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b556acbf64c05b65f06359a312ff43340bef2ec8516167b6d80b1a06f85d2bfa"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 00ed7b4dd543 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bf08e330d12c0366bdeabcb9d81fd6d82b608e25b5239968c77d81f4ecd24979"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b901d07b219a42b681ebb1cdaf91fd67a879cc67351067a0d34680d86058e611"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e119380c0308 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-f4b30371de164ff0b454124f7a34855f795333b948b81691279946259d4ef250"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

<a id="canonical-209a468c80d4fce761c18d16ce647ab1ad7bc34147015ff13c595e534a741df9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e119380c0308 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3635a1a41bfc07f2210bd8f7c2fe0a13b4f49d30cd8b81c53bc6c630c13d2cb5"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e119380c0308 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0eeb0d8edffdf22a4fec11130fc4f1b70414734e4fda7b12481b7f9f11baa6b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f2946c4e86f3c1741712e882a1e60d976932bcfcb3e7f415690ea6a3edddc9ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-423df6922947b54827f69233c83a9757d4d291f5a95f84db4f28a072c1d665b3"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 123386411cd7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-d65f704009ba2fc50e55cfff9799d972fee56402c76cc2f84ec9bf6f52110938"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

<a id="canonical-db5a626a8ae9b4801164813f7b5a50873aef45b576513f77b84c03d72644aaee"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 123386411cd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef52e453f146ac7b62489e4f4bd03a039326f7929c339fd73caea63b2128bd82"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 123386411cd7 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-45936111480a3a4be47fa0c5096d7c0a50f14cf1d47d0c272fc14753a6e753bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c6b147922806def52551556ac91b348d0205c9373ea2a4ec689be97deecb81"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e91b507f4bdf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-7de149b5ceed58b2572e9aceb64afabe0ebc401b3e7ef523b3113b0fc4f84784"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

<a id="canonical-bb88bd7167f3d9b1a8ac3c5368210cdcf2128d8bd77d786e8d07b44d7f637ec0"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e91b507f4bdf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-236b5e1cad236e856d8f88457400254cb8d2b740ee5ba9428e1f146dc88a4699"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e91b507f4bdf / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-907f50f864e14ec5a38d72f62a0e863830ba537307790091d05fee3d4baef8ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01d8dfa28ef497e641a04cdad13a8049442185975673c8628173afa9356de8d8"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f635aaf37fc6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3f16a8231b24af8d3d8a4df6773f42a3f776ee26598568a3b7da422c00eec5c5"></a>

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
disable_path_normalize = {}
```

<a id="canonical-d278853d8f4d1636edb2fcb7ef652817fd60ca040c902e9603725909ddd7e0e0"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f635aaf37fc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08a2c389d824af36a763f980a7a6e28db557bc4c109947828132d035e3de420c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f635aaf37fc6 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-07a49a97b563f697aaf5914f6252ede0564f603ba2edcdd289c645d2b380d260"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-133b35a53296d8ac5b76334a31cba0e567d8e677c729a4c7fcd8830180d82a3a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / de39fbe279be / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-11a2bb54dc341af61220dbd21db2818ad0f8dcad30ddab4bd06c923f8970fbd1"></a>

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
enable_path_normalize = {}
```

<a id="canonical-d5ed1dfaeb2c58b4eecaa58ca3ee4308b066387bc0a16fc922d12c22b4e09448"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / de39fbe279be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8162836a49705fdc4ce3a67a9875570c32264a84752d33dfe3ec2e7241720e78"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / de39fbe279be / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e25fd1be9050e8076140f36ad748b8eca255785d47d2e0d14707da9d1d37433"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d64d2f59362d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-5b4d52050c0ec62faa3a7c8a0824782d6f96c13a853846034de048bddf89cbc7"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-89916c31046d40304b29a13ca4a3e50a73522d1caa756ffea3fa2221f8c5a439"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d64d2f59362d / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-021.md#canonical-d69b46bd4831855f97bd22d872ef51dc0a84ef0d603a9781a9e818a5158a743b): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-021.md#canonical-e7da24af1ba785def044b499a62b8864e8a66e708fe9e4bfd7d38dbc1ebc2b98): complete subsection reference.

<a id="canonical-4213921de526f8fe6ac2444d46c7bd60bdf05647e391a7f4f8e39b45b88639c1"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d64d2f59362d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-021.md#canonical-d69b46bd4831855f97bd22d872ef51dc0a84ef0d603a9781a9e818a5158a743b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-021.md#canonical-e7da24af1ba785def044b499a62b8864e8a66e708fe9e4bfd7d38dbc1ebc2b98)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcdd7d1610b6df9b9408cbd73742a31f7ea262cf5c84dad7cbefac9859007ccf"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c3fb431255c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-83004dda0d772f14c734364498c0d066fc8911b0fdc10f4520177860ae24fa6d"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-a018e54fbc09bc071301c1244465f6946c47513e92534a1f1c47307953701d2d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c3fb431255c / 3

- [header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5): complete subsection reference.

<a id="canonical-fab96a1db02b5d2a0649dcb1dd4fcda7d14f1576557bbd192236fb6e935e8140"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c3fb431255c / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6990df5596ec2a0be3ca9479ec20f7831f4723d4feced1451c809a99c3709dc5"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / da213207d1ec / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-768a28ba3af04bdc2ade18fc2da448a9b454369c8f401b8c9b7b557a1d9a399b"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-c0aada45d7446d5ba38d9f6962fc449d301578df3dacfc35206fd9ba639480ce"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / da213207d1ec / 3

- [default_header_transformation](resources--workload--reference--group-021.md#canonical-5be5dd30157944721072ae59232872c29fa9aea86107d3d7520f515be23d5c53): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-021.md#canonical-4cf277354f3cd3e85561a1ed834b9f42ea55afe1f7906f34222880fe84ff13e1): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-021.md#canonical-15ae327d09b5837fecd3ca7168f55109264965e3a74e3b3ca43efa20f2721036): complete subsection reference.

<a id="canonical-328bb1ba8f199da494dcdfb72c8621bf93052179d0fd3042866fdf4eb896e41a"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / da213207d1ec / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-021.md#canonical-5be5dd30157944721072ae59232872c29fa9aea86107d3d7520f515be23d5c53)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-021.md#canonical-4cf277354f3cd3e85561a1ed834b9f42ea55afe1f7906f34222880fe84ff13e1)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-021.md#canonical-15ae327d09b5837fecd3ca7168f55109264965e3a74e3b3ca43efa20f2721036)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5be5dd30157944721072ae59232872c29fa9aea86107d3d7520f515be23d5c53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3435ff6f67256dc615cdeb10584b28b02018c283df19820960d0d254b0d1c9a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / dd74ec0abfae / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-897fce119d234ebb0d0089c3a64703bbb87721e15336855ef29704e226cd86e7"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

<a id="canonical-facf9001a99cc86cd5295c752e659b82d4cbb41821f6a92a44a9aadfee20a268"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / dd74ec0abfae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b43abde00bc98facaa83b07825b04af1e59515ba9dd82b7ee371be11e3aaedab"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / dd74ec0abfae / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4cf277354f3cd3e85561a1ed834b9f42ea55afe1f7906f34222880fe84ff13e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff2f0e11a4f31459f1b690515d8b9e1c74d3e399814d6680386430cb116e7248"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 54a804a0f93d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2ed5614739d57f9ebe97662a11f8c09fffdbf906a4bc7eb1b5db4b8e6fadd19a"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

<a id="canonical-6af935c6d43af0289634b56e8e07bcc0beec1a45e3d21093dd6f00810faf65cd"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 54a804a0f93d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9e6b974c0e7f17f8d5ad98be34e1ecf15632166e92559b9331bdf7c3cb416d6"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 54a804a0f93d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-15ae327d09b5837fecd3ca7168f55109264965e3a74e3b3ca43efa20f2721036"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be89ba869b24243cf6f5b3b980cce33e87d8df30a46b3d1fb3f33ee0020e7fe0"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ac61e3f5669 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-021.md#canonical-72bb550fa0704dc5fdbb634216dce12c47ec40c67440b9f0f687c40de494fc5b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-ce2e3b29879828821302cf11ac85a49ca0a9c9148c71e1e57e348b8a36a91b87"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

<a id="canonical-47eaddd6d4b440bb15518687806d1e88eabd2ccc75996913f1a14c06156c384b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ac61e3f5669 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a20666a580385a8ad7aaf47ed2628a8ccbe4e3f6c3e93ba3060155d449ff3978"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3ac61e3f5669 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-021.md#canonical-13b5459fe7713ac8bbdd3e30169362a77a69d9411b807d41a8efd0cd9ede67d5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d69b46bd4831855f97bd22d872ef51dc0a84ef0d603a9781a9e818a5158a743b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a1fd52feaff7866d1352a8a29ba4657cb619ee3c01264a8ce408bdb8c3a4a42"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c6fe21aeb83f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-45ac0bd2f032f81c5e1c2dc2aac3fce6d6f1f2f36fe2b2484bb8930c55adf3bf"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-07d78beaf88c03ec9430ac158154d77b7f7498c80c5bdb83ad011a41fb75387a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c6fe21aeb83f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c38f3edfcad989f7c5d231d9f8becfbdb0809287b0ee1f892833af71a4beecb"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c6fe21aeb83f / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e7da24af1ba785def044b499a62b8864e8a66e708fe9e4bfd7d38dbc1ebc2b98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2563571400d23b094693f3fba0a4fa398ff3bd092cd0be129a9e0abd3360bb8a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1a3f5911284a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-57e7d7e6d0a5bf308c80a94da0a18c410d2264096b435979f998a506fbf81d25"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-04edb1e1d537183431807b0241a5cb17fc1ac6e145ddeef54b682695f5707147"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1a3f5911284a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2926e63eed57aedf049c07edb806e603c49a28fd45923354732e4a09e36e6be"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1a3f5911284a / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-021.md#canonical-85c4a547dd87ab18a6b5ba840120febfa0334e8d1594f134c989cf3d85454f50)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-65d5ef0a35d63014f59cc3c1f97bdd4784c7ff7a0fa9a69c031410d774bb9923"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0c9e271e40496a102bb00cfb0fb7cfb565b96a47344313954cf9a04f6fac68e"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7333bdf9bbb3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-f2713fcf0515d679dabf67894387ce1c83ff590651639945a07ac68968905174"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-2df8961fa6b8b5cc552b85dbb3eaa41008a101bb4110289a94862a2ac3b5eab9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7333bdf9bbb3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b76abc45990a4bc1c5fbbaddb5d43acc8865e9ef9f63b5038fb5738700262c78"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7333bdf9bbb3 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b31dda061c4a7caa33f3654faca7ef2c566e9159cda457e0250aef367bbb2dca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b2c3f48462ea067eddbb9e1f3b5c8ed972b7d910beab8e911b72edb744af96f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6cf7f20f75c0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through

<a id="canonical-848ce3a899e8a601ee5f7c264c688b664c84c5205823f1c049ce90e0d6ff322e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-ddc9f857047f5198cb1a40cb017bc38c39899b4f36387eb6e43e8016e2e39b55"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6cf7f20f75c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3de227685dc6164916c32d2250feaf435e675c6426c2a7900d9b15fd6c0c7136"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6cf7f20f75c0 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31de6d73bfb660debb8221b338319afe105f288b2a1f8ca1d755998fa95669fe"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 33cadaf7e6ca / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-9c3dc7ba1b8112624646a21d2152552987ab42593625cd89db39cd750e6a9482"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-f832d92878aad0ea18e96dbf05e6d543429136f09c2088c85c4a0d2b0fd240ef"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 33cadaf7e6ca / 3

- [certificates](resources--workload--reference--group-021.md#canonical-f7d4f4e0d4bfa9a7a0db8c94cb8262b9ec7c9f17a19a2f94883638704d97c553): complete subsection reference.

- [no_mtls](resources--workload--reference--group-021.md#canonical-964fa2b6855e66624403966ac7c60d60970757f1c25f955914d916517a700736): complete subsection reference.

- [tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628): complete subsection reference.

- [use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1): complete subsection reference.

<a id="canonical-7d892a36e835924f73378a6d2d68adc6fb2850cd570f00a75f56306bbe82e82a"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 33cadaf7e6ca / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-021.md#canonical-f7d4f4e0d4bfa9a7a0db8c94cb8262b9ec7c9f17a19a2f94883638704d97c553)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-021.md#canonical-964fa2b6855e66624403966ac7c60d60970757f1c25f955914d916517a700736)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f7d4f4e0d4bfa9a7a0db8c94cb8262b9ec7c9f17a19a2f94883638704d97c553"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c7306897b098495bb73ff9c7670fb64a6c98345b79776baba3fb3a711f9e263"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b78f11267cfb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-c21530d189d15db0db3b78b0373f7d8adad430489713275abc2d9da7f5cc3e5e"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-8cd6e5d41a1dcd603a7756d95fc5caaad84690037d7fde8ceeebd83cc6af263f"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b78f11267cfb / 3

<a id="canonical-1fb9626deb7839c106816966bb80c6a9915a6de351cd17c578d2f16bdf0c4fa1"></a>

<a id="canonical-cbdc6d7b2c37b7b7c63ca5d9b3ffdef1ac15e16b54c4293e1137492a3e116fb7"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b78f11267cfb / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-b2c880c5dd6cee4c31d0a59d7b4ea053f07244ee5e2edb62b1e8ffcd4ce6ebaa"></a>

<a id="canonical-0df07037a7a8b48132ac779b5b64f032f4f17b0e6e7165d300ecb55bddd71a29"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b78f11267cfb / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-aaa9d55766b0ccf0f3af35e4a4ef45fb1cfc5a2690ef019adf887bc0d5ad11c3"></a>

<a id="canonical-5535797b36e8f0daefbbdaa9e3a417a0e1b8c9468f9d9b37b9a3650d0aa28b5b"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b78f11267cfb / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "byteLength": {
      "max": 64
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-9c3de176be74f2ed5affbf94af377c3840edda8b59df6cbc981aa5a79b14f30b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b78f11267cfb / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-964fa2b6855e66624403966ac7c60d60970757f1c25f955914d916517a700736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-364d737feeace6b77c69c1e0ca0555d5d9a1c7dd032a1cbfd3cf250cd4f337ec"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 43e1e93d8db7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-1d00838c3bbdb0931751b5d820a37309f39527849ef3d384cc7ada05779fe42a"></a>

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
no_mtls = {}
```

<a id="canonical-724f840332aadf1db6b4a19e2735984bf51df9416c0dddad482b219b95981aae"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 43e1e93d8db7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b092d4ef7f84d618eabf2eccd21b7cbc95f584053971f3b82d9f79bcf732833"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 43e1e93d8db7 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-984bd68f76bc21fbaeeed9b849fe17027ccc23210ff73105cb52fe2aa076e522"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 06611f8e121b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-7325e58cb7d60670838c72f1bf610a9e529c015b783dffa4da3a6d3303bf6993"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3eb23ee0188a60ad80748df2fdf647fbbf8a7cb861bb9f98e15940b9bcffb2f8"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 06611f8e121b / 3

- [custom_security](resources--workload--reference--group-021.md#canonical-9219014a527a48eba24b3118e835408bfaa66d5a144502446ef2cfc62eb30a75): complete subsection reference.

- [default_security](resources--workload--reference--group-021.md#canonical-c93c0d63277e79e04b9e1864a2461119ad15a6c7b91831103ed7944474d45364): complete subsection reference.

- [low_security](resources--workload--reference--group-021.md#canonical-86e0133e716c27bdca0da7b156f7242cdf9f4299e1c7fa6cb70eaaa9ebd51ae2): complete subsection reference.

- [medium_security](resources--workload--reference--group-021.md#canonical-0b7953bfc5820a3da6192ab9491a1a626e39a22d43983ec2c513e30e32378b83): complete subsection reference.

<a id="canonical-8d9bf0bfa75eb9d6223d8abcf6c62d823464e4df4805b50f2d0d88b65ab93b50"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 06611f8e121b / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-021.md#canonical-9219014a527a48eba24b3118e835408bfaa66d5a144502446ef2cfc62eb30a75)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-021.md#canonical-c93c0d63277e79e04b9e1864a2461119ad15a6c7b91831103ed7944474d45364)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-021.md#canonical-86e0133e716c27bdca0da7b156f7242cdf9f4299e1c7fa6cb70eaaa9ebd51ae2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-021.md#canonical-0b7953bfc5820a3da6192ab9491a1a626e39a22d43983ec2c513e30e32378b83)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9219014a527a48eba24b3118e835408bfaa66d5a144502446ef2cfc62eb30a75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37d2180e201ff28d3c45be7233a75ed95d52126f81e07c7f36ecdbe26122899d"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1996fc585168 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-42e3d7eefafbc135818518fd5795c2970beb02df7d75225bddd9bb32ecb82432"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
```

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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-f99d75e14575ba4e3835eba09c573d684a34bb164358440c0308b84d00555016"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1996fc585168 / 3

<a id="canonical-5ae976adbff6d2429354b061770b6bcc55c04d8c769f1e5f4888cbe94ab5601e"></a>

<a id="canonical-1e45fbb5c49e426b246216cd1e9c9ab2386eaa584878a4256295ed9bff6eb45b"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1996fc585168 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cce03b04650319c2f43bfe05e7af91d4b4ae9cef03b67652ecd572abc908d1d1"></a>

<a id="canonical-36328b4d63f7cae973026f78d6966830d9052f128d5505f50b5387e88455f924"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1996fc585168 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d1e510d680d7fae17c48ca149649fca50d3ffd4eedbd639b74c5a088c72028f5"></a>

<a id="canonical-a34de990cf1094311f71809fd826b1aed67f744827bc2c69dd4d6744dc126778"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1996fc585168 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-74bdacf496dd6554aecdc1481c4e73d41d32bd19bea8b0f748c1f914501c643f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1996fc585168 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c93c0d63277e79e04b9e1864a2461119ad15a6c7b91831103ed7944474d45364"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7df4365e290c4e028bc55dda196d5de5b15610683d0304173f93735232d18dd"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7904620fc0f9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-ed729fdc72af725fe894b92fb0f62aacf216f6efa2eb63e5d2eec91875b22f40"></a>

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
default_security = {}
```

<a id="canonical-a0a3ebf19a5482bf730230840fa03c5fe03ef8b9cfe48495b0ec85d970deb594"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7904620fc0f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57a4ef6dd3ffc26a9fc2fa07de5892c1615d5401aef56cb475e9cc8fe29afaf6"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7904620fc0f9 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-86e0133e716c27bdca0da7b156f7242cdf9f4299e1c7fa6cb70eaaa9ebd51ae2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3a30d7dd7615c8e657b012a4b5eb22aa42811299abbcad039425114b1a5d468"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / cc674d38ee0d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-4f13fcd13084b6b931da6cafd215239c9c1d74df751a9775d9d591e7f792f3a9"></a>

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
low_security = {}
```

<a id="canonical-cb872d7260f2d4f1c66c290bd232e7175b267eb052003bf0afc122fc42377048"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / cc674d38ee0d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07d27ce66c241279768ed56badcf82cccdb3ce6d9e01c7b872e8b52cc485376e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / cc674d38ee0d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0b7953bfc5820a3da6192ab9491a1a626e39a22d43983ec2c513e30e32378b83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0caa2ebabc8557cf5e8183c17e5955f507cb465b6c0fb7b312ddbd0aa9f71c97"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e333070c0e22 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-179db5c6a1905327264c97e9c2782d7e33f9ed6e1dc1dbd585ac9d886ee08e4a"></a>

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
medium_security = {}
```

<a id="canonical-8d14d867aea9657585644844a95df999b86e45c05536a6c6c77589f0326abf9a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e333070c0e22 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b486f535fa5a1710a98e53a1632535803c4068657bdca860176a41958753ddd"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e333070c0e22 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-021.md#canonical-6b5b9d4ff538cc3371b6209dedd32e4269993255bbe3624913a11f7408a1e628)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-727c1563aa63f00a32e4251391a1c825bd458fa3b8fbd4faaecc2787565ac901"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a3c6308a623c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-e667a6fe558261a872a8360b5cedbac32cc2281c0248c81f9a86be0bb062e9ec"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-190ab254e7cec478e8194564a41e98b6bdd212ca45d6ba33edc74e4c1a66f53b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a3c6308a623c / 3

<a id="canonical-db451fc67a29d78a18bbad9682088fa691772f1e8b199549d969313f27b291a0"></a>

<a id="canonical-45f5014db54ca43bf821251f5e7c52ab1de8fa757719975231e2800c9e99e0d6"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a3c6308a623c / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--workload--reference--group-021.md#canonical-81af31a068b3c662465d9f0f7d5e183ec0b185a6f1c526dc205a975ab84a2c4f): complete subsection reference.

- [no_crl](resources--workload--reference--group-022.md#canonical-1df4847bf6be57a4006e0010cae62f49cbd97b40062804ae95c7976517a39402): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-022.md#canonical-34bdf188ecf8ecfaf9609b4979bffb1166a5e4cb989a0e58e8b47c383e2ac65f): complete subsection reference.

<a id="canonical-07c54c3fc068a67ef4596396fe2ea560867db47be4b4af718267042c8c66dc24"></a>

<a id="canonical-703d86cf5c1526d1f038bfadd3967cb6bda801e10379c07e911948757cebd9da"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a3c6308a623c / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--workload--reference--group-022.md#canonical-1fac8e48a8f05439c39a9dc94fda3a5a57e1d718c809df562b8a69fef74f0913): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-022.md#canonical-8dfda5cf4710875fa5ec033f08a74d62d3765c9d97905d178c4568e298951e58): complete subsection reference.

<a id="canonical-257e0ea11b8b785fdadf845b2da6550f6ea553b55e4321014ed91a26c8f15724"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / a3c6308a623c / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-021.md#canonical-81af31a068b3c662465d9f0f7d5e183ec0b185a6f1c526dc205a975ab84a2c4f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-022.md#canonical-1df4847bf6be57a4006e0010cae62f49cbd97b40062804ae95c7976517a39402)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-022.md#canonical-34bdf188ecf8ecfaf9609b4979bffb1166a5e4cb989a0e58e8b47c383e2ac65f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-022.md#canonical-1fac8e48a8f05439c39a9dc94fda3a5a57e1d718c809df562b8a69fef74f0913)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-022.md#canonical-8dfda5cf4710875fa5ec033f08a74d62d3765c9d97905d178c4568e298951e58)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-81af31a068b3c662465d9f0f7d5e183ec0b185a6f1c526dc205a975ab84a2c4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
