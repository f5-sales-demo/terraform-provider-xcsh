---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-b3b04e9c813aa175e6ef0df958f74179887ae82315103ebd207f6f05d660a8de"></a>

## max_cookie_count_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 4

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](resources--service_policy--reference--group-003.md#canonical-0cab19be06c04aa2154b0260abc5ee9b5fe052b93f4aba147e6511ad3b2a512f): complete subsection reference.

<a id="canonical-3fefcbe59d253f67c10c83698d1156a1b7adfad37e17c031ab755f0012f87ceb"></a>

<a id="canonical-d1bb8e8a33167a7388cba4b15f4dcbf08568fc6bf29e5526a103c044471e7e76"></a>

## max_cookie_key_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](resources--service_policy--reference--group-003.md#canonical-57e45b6a86a0bf0f06fc6a50fa7ed2e321034634b1520e18cabae726cfcdba86): complete subsection reference.

<a id="canonical-1c846dc009d5ede0152c80167ede695bfd043d2d0dfd702ba91b2006908145e1"></a>

<a id="canonical-6774036f1d84d108ea0624bc40dea4ef64c697838b61da9f1ec421dbd9efc742"></a>

## max_cookie_value_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 6

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](resources--service_policy--reference--group-003.md#canonical-839732bbfeac4fb9b1a8c22f49f369f4054730f012e38dbe52a35670172faf92): complete subsection reference.

<a id="canonical-5690aa69b7bc21a9122a7b6da36ee187272698751e7c7d090f2dbec9b34308fd"></a>

<a id="canonical-e3b36512ca4c6ffb54a49286fc1a55b956493a92411053ee2e8a75344acf1c05"></a>

## max_header_count_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 7

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](resources--service_policy--reference--group-003.md#canonical-46d80c5b2898414e3b4f722357a0046a99975a1432bb2d2a80bf945d1c95c010): complete subsection reference.

<a id="canonical-1740941ace534eb5ef460a6469572e60cf157622921537e7bae62e22f5cb36b7"></a>

<a id="canonical-3473f302f24e166df74a57891528cbc53d0c03a532d303558564be24e36022d9"></a>

## max_header_key_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 8

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](resources--service_policy--reference--group-003.md#canonical-c2f6db064212244c6a12ef87a2182849564125368e0994fea551d34a32e6a895): complete subsection reference.

<a id="canonical-552e813cb27493da2643a71128ce6c43a8945a3c541c411a55a354d7fd9695d0"></a>

<a id="canonical-96c24f82f10e079318b64182228e95e5d1e6eddc8a149db1e4a7c5a5b1c75ffe"></a>

## max_header_value_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 9

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 64000),
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
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](resources--service_policy--reference--group-003.md#canonical-259270519f87ba85b8ad58f2c90ba43e652f49019f62e33a03c693d7dfce07a3): complete subsection reference.

<a id="canonical-e74663534f076e8095e5a5540e89bfbb133999b73b5a06831fcaaba76a876a1c"></a>

<a id="canonical-f4d09126fdf18b477e8ddd9334dd100def8ffc5af618448ce444cb98cd606875"></a>

## max_parameter_count_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 10

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](resources--service_policy--reference--group-003.md#canonical-d1cac10a0737651f396687cab5db25bafd2dcfb15fe319aae9e39d718d34b7d7): complete subsection reference.

<a id="canonical-f1a95629f30c789bc9181ccc2d2f59424a3fb241cc0906f8cc61cf95cc55a552"></a>

<a id="canonical-f778ac189834b3785a927ad1be8eff354db622edbaffe771f868dccdc3b35679"></a>

## max_parameter_name_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 11

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](resources--service_policy--reference--group-003.md#canonical-02c5016d9ad57f1935a6279072fe9ca3ac54258e483d39ffde101bb33261760e): complete subsection reference.

<a id="canonical-7ebd0879f9618e74440341cb5e3db8547bba20f0004f29864656311874bfcf9b"></a>

<a id="canonical-0e59af96c5e51f8deca3e3222dbeb9fc668657c2a509df37e87ed90bf0ec811b"></a>

## max_parameter_value_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 12

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1073741824),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](resources--service_policy--reference--group-003.md#canonical-e0ee6bf90f172bd9a3c8a7a6586a4df9bf9d831a8801b2b8686418f4e2c38461): complete subsection reference.

<a id="canonical-b78e3d3f096503d8a4e2810e2d4e9ff7f22eeec51562d3e5a388babee428a217"></a>

<a id="canonical-1071dd4148501c35c2ad76fcb5d03017f85f34f3bf5071d7b77b9a198604bb5b"></a>

## max_query_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 13

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](resources--service_policy--reference--group-003.md#canonical-4722a54e9bc90da3eecb884f5679625af561d20842785d0b905d5f2936c71812): complete subsection reference.

<a id="canonical-55fc86eb832c7077737fb934a71a2ea34b391c4d2aa08f35324991bd75e30c3d"></a>

<a id="canonical-28a4ea3a82b711403411a7a449f7bc6b9553116032946ae7aa387a205dffb66f"></a>

## max_request_line_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 14

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](resources--service_policy--reference--group-003.md#canonical-768fe702044fd2e8ec77981ff2f8c3ee9ce009beb4c8ef975a9466cccdbda56c): complete subsection reference.

<a id="canonical-cf7c4ba22ee42a28164bffec1950058750f29ec82ddbb08407002d3ef306550d"></a>

<a id="canonical-c3a3614cdd05a607e604f4a7b9216c4fc8473dadefe6bf7e77de4c439e62ba55"></a>

## max_request_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 15

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](resources--service_policy--reference--group-003.md#canonical-a50e03638c012d15c4d441e37480a219bd4919b22decc67f4701d71a581a424c): complete subsection reference.

<a id="canonical-d621245f1d3f7a3994332a8b06552b822ab9eeb64a202f22ab38e3291073e1e7"></a>

<a id="canonical-88da5aadb5424587000dd34239875596ef7dcf9a42bc8182b55e3043139b9387"></a>

## max_url_size_exceeds property — rule_list.rules.spec.request_constraints / e35c31160092 / 16

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 128000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](resources--service_policy--reference--group-003.md#canonical-01586884d6ee5e484436de728e84d2020edda5aede22a45e1310d6ded48e6e1c): complete subsection reference.

<a id="canonical-cc0e0d7cb5d6b168743e0e8c1c7a2c018552b879d48cb7887b69263eb3a6972e"></a>

## Next pages — rule_list.rules.spec.request_constraints / e35c31160092 / 17

- [rule_list.rules.spec.request_constraints.max_cookie_count_none](resources--service_policy--reference--group-003.md#canonical-0cab19be06c04aa2154b0260abc5ee9b5fe052b93f4aba147e6511ad3b2a512f)
- [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](resources--service_policy--reference--group-003.md#canonical-57e45b6a86a0bf0f06fc6a50fa7ed2e321034634b1520e18cabae726cfcdba86)
- [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](resources--service_policy--reference--group-003.md#canonical-839732bbfeac4fb9b1a8c22f49f369f4054730f012e38dbe52a35670172faf92)
- [rule_list.rules.spec.request_constraints.max_header_count_none](resources--service_policy--reference--group-003.md#canonical-46d80c5b2898414e3b4f722357a0046a99975a1432bb2d2a80bf945d1c95c010)
- [rule_list.rules.spec.request_constraints.max_header_key_size_none](resources--service_policy--reference--group-003.md#canonical-c2f6db064212244c6a12ef87a2182849564125368e0994fea551d34a32e6a895)
- [rule_list.rules.spec.request_constraints.max_header_value_size_none](resources--service_policy--reference--group-003.md#canonical-259270519f87ba85b8ad58f2c90ba43e652f49019f62e33a03c693d7dfce07a3)
- [rule_list.rules.spec.request_constraints.max_parameter_count_none](resources--service_policy--reference--group-003.md#canonical-d1cac10a0737651f396687cab5db25bafd2dcfb15fe319aae9e39d718d34b7d7)
- [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](resources--service_policy--reference--group-003.md#canonical-02c5016d9ad57f1935a6279072fe9ca3ac54258e483d39ffde101bb33261760e)
- [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](resources--service_policy--reference--group-003.md#canonical-e0ee6bf90f172bd9a3c8a7a6586a4df9bf9d831a8801b2b8686418f4e2c38461)
- [rule_list.rules.spec.request_constraints.max_query_size_none](resources--service_policy--reference--group-003.md#canonical-4722a54e9bc90da3eecb884f5679625af561d20842785d0b905d5f2936c71812)
- [rule_list.rules.spec.request_constraints.max_request_line_size_none](resources--service_policy--reference--group-003.md#canonical-768fe702044fd2e8ec77981ff2f8c3ee9ce009beb4c8ef975a9466cccdbda56c)
- [rule_list.rules.spec.request_constraints.max_request_size_none](resources--service_policy--reference--group-003.md#canonical-a50e03638c012d15c4d441e37480a219bd4919b22decc67f4701d71a581a424c)
- [rule_list.rules.spec.request_constraints.max_url_size_none](resources--service_policy--reference--group-003.md#canonical-01586884d6ee5e484436de728e84d2020edda5aede22a45e1310d6ded48e6e1c)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-0cab19be06c04aa2154b0260abc5ee9b5fe052b93f4aba147e6511ad3b2a512f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-deabe49d43f291b8682180c279f284e0e89889a1dbd34d5ef8a2204e64917177"></a>

## rule_list.rules.spec.request_constraints.max_cookie_count_none — rule_list.rules.spec.request_constraints.max_cookie_count_none / 8eaf66890f99 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_cookie_count_none

<a id="canonical-8ff641e27180f2d4cf2e4be032523a75685c9cc7285a087b627eee0ae1c5f35b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie count none.

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
max_cookie_count_none = {}
```

<a id="canonical-28fffcb8189f09c8a52571ae5d92d1c2e038361c270b4997ee92f4bd603088ec"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_cookie_count_none / 8eaf66890f99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90b1593fc6d3e6ec0c80d2f9fdeca29856e9344f1e17f80dc092dc03a25c759d"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_cookie_count_none / 8eaf66890f99 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-57e45b6a86a0bf0f06fc6a50fa7ed2e321034634b1520e18cabae726cfcdba86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e739a506ef5c89f8cbb742d593c8c0c5096b3478cb1ab12ffab0606d8c9197"></a>

## rule_list.rules.spec.request_constraints.max_cookie_key_size_none — rule_list.rules.spec.request_constraints.max_cookie_key_size_none / 47d3a4192a40 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_cookie_key_size_none

<a id="canonical-7aca03db5d566c338ebc70db39442ba295c86e96fad74d043a199af3e5628013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie key size none.

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
max_cookie_key_size_none = {}
```

<a id="canonical-57681991fec37de280e8cddf78c3ff29e90fa4653ff1204f8a34caebebbfe924"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_cookie_key_size_none / 47d3a4192a40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19cd619cee0218797487aa1068803bb9cf8827823c67dce84b0188b49feaf65a"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_cookie_key_size_none / 47d3a4192a40 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-839732bbfeac4fb9b1a8c22f49f369f4054730f012e38dbe52a35670172faf92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16b63a40636969e905e6fc88bc2f7e80b999eb7f7d11a1c5f23c7a1444e44212"></a>

## rule_list.rules.spec.request_constraints.max_cookie_value_size_none — rule_list.rules.spec.request_constraints.max_cookie_value_size_none / c4e524845114 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="canonical-d0591c0821a7158ab2f06a1ea558bfed9fd2f70e0d862724868b0d35240ec04b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

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
max_cookie_value_size_none = {}
```

<a id="canonical-068570b4752ba8ba97c80e590cf6f137f54007b13eff359b43164d1bdfddc065"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_cookie_value_size_none / c4e524845114 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3883e1e46143db06da593e8aab17f41e4eb52eaa834b9c0855f4df415c5c3acd"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_cookie_value_size_none / c4e524845114 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-46d80c5b2898414e3b4f722357a0046a99975a1432bb2d2a80bf945d1c95c010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ed9468a9653a249f9aca186eb25d871d7b2fbd0b776675af6ac7b9a01a56ff8"></a>

## rule_list.rules.spec.request_constraints.max_header_count_none — rule_list.rules.spec.request_constraints.max_header_count_none / 0c96b0338e5c / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_header_count_none

<a id="canonical-ed77a73aba6179c2fbf790d2ef01cd8ef256a94a00e8f3d4df3df3033ce3dea7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

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
max_header_count_none = {}
```

<a id="canonical-76d01fdf13b502db79c19e5217e530c6cb4173cfeef9963b36deb63e318bc26b"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_header_count_none / 0c96b0338e5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9558660ac8e5cac7643b6d3d157830c68281c0a115029bc03538674298c0336f"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_header_count_none / 0c96b0338e5c / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-c2f6db064212244c6a12ef87a2182849564125368e0994fea551d34a32e6a895"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f31d57f7131512ed21860a95af7fb78d0347e4b920af84019732ee94284d4e1f"></a>

## rule_list.rules.spec.request_constraints.max_header_key_size_none — rule_list.rules.spec.request_constraints.max_header_key_size_none / 7bfc507a0447 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_header_key_size_none

<a id="canonical-b5e6aa65a21c75df647b982abfdcf39d663b3c7f63e08d0d0358fd8d42fa9529"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header key size none.

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
max_header_key_size_none = {}
```

<a id="canonical-23f4ff28e0fefe42fb61045e34edbf5d975b46d38f80f6b7906b08fed30b0d12"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_header_key_size_none / 7bfc507a0447 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf0f3d4eca98fb190a71a61a144f6b206852147446edd3fd7119aff2e925ab54"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_header_key_size_none / 7bfc507a0447 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-259270519f87ba85b8ad58f2c90ba43e652f49019f62e33a03c693d7dfce07a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-665d164848b5ab02edc63bc06e9eb62a640dcd12255a888e8730bf9f836fe757"></a>

## rule_list.rules.spec.request_constraints.max_header_value_size_none — rule_list.rules.spec.request_constraints.max_header_value_size_none / 1471aac43700 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_header_value_size_none

<a id="canonical-208b62dde92d80034a7f1b7e4e7412eb30579d8ed2f2a95815b37eb93801a0a5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header value size none.

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
max_header_value_size_none = {}
```

<a id="canonical-76258cf3fb98d01cdb112e71c321382f3dbca3cc7a1a345ea73ca2303a512dda"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_header_value_size_none / 1471aac43700 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aee4cfd27ac5a44b4076135884fd2fe7cae396b61c6cab703e33ed6f86ee0065"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_header_value_size_none / 1471aac43700 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-d1cac10a0737651f396687cab5db25bafd2dcfb15fe319aae9e39d718d34b7d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2afc876a94ad242c8a1a8c520e42b8b29c0ee7082bf1c10c0493d55c9e491c50"></a>

## rule_list.rules.spec.request_constraints.max_parameter_count_none — rule_list.rules.spec.request_constraints.max_parameter_count_none / deb4318f51b0 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_parameter_count_none

<a id="canonical-450cfa03ad827a9671ef4ecca111cf74e26bf0c080d4e272a78c53f4be90efca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter count none.

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
max_parameter_count_none = {}
```

<a id="canonical-4b222f59bacb5ba8d8901ef88053f54c65a3b8aea1c3c84ea05beab222efc4da"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_parameter_count_none / deb4318f51b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62989d7a3b9007992cfd547ab42aa69afe8751812725307c97251474114b278c"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_parameter_count_none / deb4318f51b0 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-02c5016d9ad57f1935a6279072fe9ca3ac54258e483d39ffde101bb33261760e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f565d8fe9a0136cc97f2c54da6a682c2a2b195bf177cf58c4264b45ea0b6c035"></a>

## rule_list.rules.spec.request_constraints.max_parameter_name_size_none — rule_list.rules.spec.request_constraints.max_parameter_name_size_none / 89d15ae8c0fe / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_parameter_name_size_none

<a id="canonical-4283d8d3694530ededa7d23174ec460dfdb28fd523073728f4c29d3c10d3fe3e"></a>

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
max_parameter_name_size_none = {}
```

<a id="canonical-8efb99e343a7c3151fd210708c7a85fe852bc6ff8b0ec0bccf72c0dbc9278511"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_parameter_name_size_none / 89d15ae8c0fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59371d87506a6d655b750ed17a3996aba8677dab75ffe2b2e0f4d40397199e3b"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_parameter_name_size_none / 89d15ae8c0fe / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-e0ee6bf90f172bd9a3c8a7a6586a4df9bf9d831a8801b2b8686418f4e2c38461"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-917b003301d354f50d1c97cba0d5b599a4a6a6f7ba9814e26b0f263d9c61ba85"></a>

## rule_list.rules.spec.request_constraints.max_parameter_value_size_none — rule_list.rules.spec.request_constraints.max_parameter_value_size_none / d08622d6ad1c / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_parameter_value_size_none

<a id="canonical-6e7d17139e61550a5c63d587b31ec61dd5fba801c24ea986783c592de906effc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter value size none.

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
max_parameter_value_size_none = {}
```

<a id="canonical-d3793b6d69a779872a562e58db66aeffb14b1a72b4c2bbb3e61bcaecefec682b"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_parameter_value_size_none / d08622d6ad1c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90fdc4789c2d5c202fab7c9660ce66c6f5bee800ba5e40c5f470a4d33f409738"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_parameter_value_size_none / d08622d6ad1c / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-4722a54e9bc90da3eecb884f5679625af561d20842785d0b905d5f2936c71812"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83c08ffb11213413f21ebe7e78d71aa70d46cc8937fa1b1eb63592990bd06d6c"></a>

## rule_list.rules.spec.request_constraints.max_query_size_none — rule_list.rules.spec.request_constraints.max_query_size_none / 2199caca9eaf / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_query_size_none

<a id="canonical-a8ae2038b75339cf45de48bc0e3a11770f774ccfb93d25af5eeef052534cbfab"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max query size none.

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
max_query_size_none = {}
```

<a id="canonical-e30438d1a03590e5685360f970e0c8509ec07fb46a4bc880ab051eec79c8a800"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_query_size_none / 2199caca9eaf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c57f70be23a5f51956ae8677874fd69ce0cde95690ba399fa608b71064c7e11"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_query_size_none / 2199caca9eaf / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-768fe702044fd2e8ec77981ff2f8c3ee9ce009beb4c8ef975a9466cccdbda56c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a7cedebdb50eb2de59bc8f0e21bfa00fea9e03fbb196d9a916f60f19935382f"></a>

## rule_list.rules.spec.request_constraints.max_request_line_size_none — rule_list.rules.spec.request_constraints.max_request_line_size_none / d6c2a991c0b0 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_request_line_size_none

<a id="canonical-7992a650ae4b203861ba01a94a8703d826afb54d46c1af01da0a453ab04e6ba3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request line size none.

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
max_request_line_size_none = {}
```

<a id="canonical-c6d9f46106409b23b4e64d6610c6591fe3a24263d558aaf886937ca95ead07a2"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_request_line_size_none / d6c2a991c0b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb4388d0ad80b561cc153741c2965aebb10ff5698ba4b724a59635afc099cd57"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_request_line_size_none / d6c2a991c0b0 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-a50e03638c012d15c4d441e37480a219bd4919b22decc67f4701d71a581a424c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48af650814eb1161fb467295e5dfd7d4de3e194e7aa3248bf6615f257d0e127a"></a>

## rule_list.rules.spec.request_constraints.max_request_size_none — rule_list.rules.spec.request_constraints.max_request_size_none / 913298141c44 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_request_size_none

<a id="canonical-63d0e8e5e0900ae2623388165ff281237416257af873074c48837e626b978f7e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request size none.

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
max_request_size_none = {}
```

<a id="canonical-07c25000f10b3d864e08fdd817bb0324c5557b0d928864334234ed415bb3a059"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_request_size_none / 913298141c44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f2dd243ad77dfbeab3729c0bc4932642bdfe487acc2a8c8a49bb77eeb41eadb"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_request_size_none / 913298141c44 / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-01586884d6ee5e484436de728e84d2020edda5aede22a45e1310d6ded48e6e1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6fbc00f9fd1945be02ae46f47affacafe36ab9ab1478882f33465036d8eacd2"></a>

## rule_list.rules.spec.request_constraints.max_url_size_none — rule_list.rules.spec.request_constraints.max_url_size_none / 7d559a24648a / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- rule_list.rules.spec.request_constraints.max_url_size_none

<a id="canonical-fd444ea04ef8ba434703cdd7b25fc9f4bfeff5cc2a0de7825a37e026d1e1c372"></a>

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
max_url_size_none = {}
```

<a id="canonical-08d62877059a0240d4aa3488370e7908a246fefb5a4c58be32bc1abef8da938f"></a>

## Direct properties — rule_list.rules.spec.request_constraints.max_url_size_none / 7d559a24648a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fab808b6cd292810d7731f372e17b547b821db3459412e890849cc8524a6db8a"></a>

## Next pages — rule_list.rules.spec.request_constraints.max_url_size_none / 7d559a24648a / 4

- [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d7215e8d8d9bdd803dbb89922558a6931b4b4a29425e4f7420c147e813c8f9c"></a>

## rule_list.rules.spec.segment_policy — rule_list.rules.spec.segment_policy / c9ee836cbe4f / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- rule_list.rules.spec.segment_policy

<a id="canonical-db18072a0934fb5a4992b0e239f6fd12340de6f0a848985db043717928d70c99"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dst_any",
    "dst_segments"),
  validators.ConflictingObjectAttributes("dst_any",
    "intra_segment"),
  validators.ConflictingObjectAttributes("dst_segments",
    "intra_segment"),
  validators.ConflictingObjectAttributes("src_any",
    "src_segments")}
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
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

Terraform syntax:

```terraform
segment_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-f40089c28307d6f8338adc8d6cd48b58aea948e6c63e0958f6a50a35b0125231"></a>

## Direct properties — rule_list.rules.spec.segment_policy / c9ee836cbe4f / 3

- [dst_any](resources--service_policy--reference--group-003.md#canonical-c2125454e9efa1d40055c176994bcb896b12c1ef043cb8d9e3703a997852ddcc): complete subsection reference.

- [dst_segments](resources--service_policy--reference--group-003.md#canonical-8b4b041004cabcfc708ae28f132dcc8e4b1b1925f873f3478fe1c45f5dff3f9e): complete subsection reference.

- [intra_segment](resources--service_policy--reference--group-003.md#canonical-e1ec829dc7c30c1c26cb7c71b0c72933890337ad1ee8bb561a5af8bf401f0ada): complete subsection reference.

- [src_any](resources--service_policy--reference--group-003.md#canonical-07ca3d6244a0367ee4b19bb6c22e30b8605cef1bea58ed9b8efc2917f64772b0): complete subsection reference.

- [src_segments](resources--service_policy--reference--group-003.md#canonical-7954f333c9f5786892f536436ac0d0c7d075d51971afe9d04e72e8684eaa5a4d): complete subsection reference.

<a id="canonical-26d89e63f9c15b19a64c1526fe0b95767a427dfab213fdb9bc52a7fd2f64e867"></a>

## Next pages — rule_list.rules.spec.segment_policy / c9ee836cbe4f / 4

- [rule_list.rules.spec.segment_policy.dst_any](resources--service_policy--reference--group-003.md#canonical-c2125454e9efa1d40055c176994bcb896b12c1ef043cb8d9e3703a997852ddcc)
- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-8b4b041004cabcfc708ae28f132dcc8e4b1b1925f873f3478fe1c45f5dff3f9e)
- [rule_list.rules.spec.segment_policy.intra_segment](resources--service_policy--reference--group-003.md#canonical-e1ec829dc7c30c1c26cb7c71b0c72933890337ad1ee8bb561a5af8bf401f0ada)
- [rule_list.rules.spec.segment_policy.src_any](resources--service_policy--reference--group-003.md#canonical-07ca3d6244a0367ee4b19bb6c22e30b8605cef1bea58ed9b8efc2917f64772b0)
- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-7954f333c9f5786892f536436ac0d0c7d075d51971afe9d04e72e8684eaa5a4d)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-c2125454e9efa1d40055c176994bcb896b12c1ef043cb8d9e3703a997852ddcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1c75bc028b6eb34db55b2c218713e09ffd7b3a80212fe70cc8a0ff703f8d4eb"></a>

## rule_list.rules.spec.segment_policy.dst_any — rule_list.rules.spec.segment_policy.dst_any / 515c4b375317 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- rule_list.rules.spec.segment_policy.dst_any

<a id="canonical-46bc78791717eb7c3251bdbbcba0293797621b715879e3e0be1e9dcd7b5fa623"></a>

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
dst_any = {}
```

<a id="canonical-1b3a46b324eeb03b7abe2a2129f46444d84c58778b2ea3572aceb32058df5c86"></a>

## Direct properties — rule_list.rules.spec.segment_policy.dst_any / 515c4b375317 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ab653989d06698acf17734ec725d5744ef121545050c56b5f01cb87c55cec51"></a>

## Next pages — rule_list.rules.spec.segment_policy.dst_any / 515c4b375317 / 4

- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-8b4b041004cabcfc708ae28f132dcc8e4b1b1925f873f3478fe1c45f5dff3f9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc7bff04146b5de84aff369e087f78b4fb9964e188b3484ffb0f885ab3d05817"></a>

## rule_list.rules.spec.segment_policy.dst_segments — rule_list.rules.spec.segment_policy.dst_segments / 0a63833d19ff / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- rule_list.rules.spec.segment_policy.dst_segments

<a id="canonical-1e80a79cf0f04eb64e6730599c576f72d2004dbb3f42143d3499bdb6d47784b2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
dst_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-1da788ea5218c691016c71f6b3f5307c8864cb7bf5c62fedb1959d4c66d9fd65"></a>

## Direct properties — rule_list.rules.spec.segment_policy.dst_segments / 0a63833d19ff / 3

- [segments](resources--service_policy--reference--group-003.md#canonical-1d42ab4f5af3e11fddaae1baf715ad73691199d675d14e503d99138ee80c95eb): complete subsection reference.

<a id="canonical-247a79126027284936ef0c38be53c9d6435d273734e508d6d90e77ef6da00b02"></a>

## Next pages — rule_list.rules.spec.segment_policy.dst_segments / 0a63833d19ff / 4

- [rule_list.rules.spec.segment_policy.dst_segments.segments](resources--service_policy--reference--group-003.md#canonical-1d42ab4f5af3e11fddaae1baf715ad73691199d675d14e503d99138ee80c95eb)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-1d42ab4f5af3e11fddaae1baf715ad73691199d675d14e503d99138ee80c95eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20aae0a0f6684af50481cd0bd79d577111712d9fb65d190d96f5a0ee87d1afd6"></a>

## rule_list.rules.spec.segment_policy.dst_segments.segments — rule_list.rules.spec.segment_policy.dst_segments.segments / 8a55844cb3b4 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-8b4b041004cabcfc708ae28f132dcc8e4b1b1925f873f3478fe1c45f5dff3f9e)
- rule_list.rules.spec.segment_policy.dst_segments.segments

<a id="canonical-be51d612a8c72d250d36b367b1c5f6abf86d11b8bb072fa84c630b29b930741b"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-ce1df686e796d778e65a99a85a238935ff0545199653b99c25eefc58ab422b8d"></a>

## Direct properties — rule_list.rules.spec.segment_policy.dst_segments.segments / 8a55844cb3b4 / 3

<a id="canonical-088fe359555b65110d1c3047759e264e3fd20c3d595d38b16aa09d4df3bd7d81"></a>

<a id="canonical-30bdc4f0eb7600137292331eb983bcbf16a931730c76c3741234f73ed55efb07"></a>

## name property — rule_list.rules.spec.segment_policy.dst_segments.segments / 8a55844cb3b4 / 4

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

<a id="canonical-1b06f80bd8e43d048af3a42d2fa302ba9587e77d24c0764a6a25c73d916fcf75"></a>

<a id="canonical-d15c1b295733d091d9c9dd42ff63d094a9de5a12e5ad0e1142324838cc8603a2"></a>

## namespace property — rule_list.rules.spec.segment_policy.dst_segments.segments / 8a55844cb3b4 / 5

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

<a id="canonical-a058872bf4bcc350d0bb43e691d38a16e5ffed94580f55f9afe78d138a572742"></a>

<a id="canonical-36fec4d4982c05000471b983217b271bc222d8432f3e6a0383aa86defb1a53fb"></a>

## tenant property — rule_list.rules.spec.segment_policy.dst_segments.segments / 8a55844cb3b4 / 6

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

<a id="canonical-a37a35261044e354bcd90867d74002d38502a42ca3b26c757bca2f087ea30abb"></a>

## Next pages — rule_list.rules.spec.segment_policy.dst_segments.segments / 8a55844cb3b4 / 7

- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-8b4b041004cabcfc708ae28f132dcc8e4b1b1925f873f3478fe1c45f5dff3f9e)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-e1ec829dc7c30c1c26cb7c71b0c72933890337ad1ee8bb561a5af8bf401f0ada"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e858712813bd5efc6c28c747592aa56ace661b6507c1a6635777e0f65a4b8ae8"></a>

## rule_list.rules.spec.segment_policy.intra_segment — rule_list.rules.spec.segment_policy.intra_segment / 023e958a05b9 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- rule_list.rules.spec.segment_policy.intra_segment

<a id="canonical-643c24a246e7218169258dad2453861a715148d868fe511c77115d892bed80f9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for intra segment.

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
intra_segment = {}
```

<a id="canonical-ff0ead638c22c23711d44622a5597b457272db86bfc7329f569f2c32f18e9fac"></a>

## Direct properties — rule_list.rules.spec.segment_policy.intra_segment / 023e958a05b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42e3b8a16bb8eb714cddbb81961c4ae0899bcd32a0f1772e892a6e55de114a1a"></a>

## Next pages — rule_list.rules.spec.segment_policy.intra_segment / 023e958a05b9 / 4

- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-07ca3d6244a0367ee4b19bb6c22e30b8605cef1bea58ed9b8efc2917f64772b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb385ff8468ca91c652346e4c030b4bd589025c92e815f3202e00132705d5a70"></a>

## rule_list.rules.spec.segment_policy.src_any — rule_list.rules.spec.segment_policy.src_any / 6c254d1ec4ec / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- rule_list.rules.spec.segment_policy.src_any

<a id="canonical-49521907b1797cfa6666abd47cc8cebe1d7020334af68937cb0bef0dbc4ae2a1"></a>

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
src_any = {}
```

<a id="canonical-b3bbf2d054ac28fb975bc9d9c18efd443ac2706ab7f3119751ab68af0dc6929a"></a>

## Direct properties — rule_list.rules.spec.segment_policy.src_any / 6c254d1ec4ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62328c371ae2f7c5724f7982a50fedcb71111c2fe111a8ed33432b6f01dc767c"></a>

## Next pages — rule_list.rules.spec.segment_policy.src_any / 6c254d1ec4ec / 4

- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-7954f333c9f5786892f536436ac0d0c7d075d51971afe9d04e72e8684eaa5a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f5c63080121f12e10d54c1eab0f4b93a645160c5a7b112741dda628731fb771"></a>

## rule_list.rules.spec.segment_policy.src_segments — rule_list.rules.spec.segment_policy.src_segments / 373dadff407e / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- rule_list.rules.spec.segment_policy.src_segments

<a id="canonical-feb6871145e3225a736f37b035cbe6259e1c559ba47cdc46a45a51430e1f4bd6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
src_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-1cf77c7b7b894ea651fa4526906355bf1fad0de4c727db2924a53a566f3bff2d"></a>

## Direct properties — rule_list.rules.spec.segment_policy.src_segments / 373dadff407e / 3

- [segments](resources--service_policy--reference--group-003.md#canonical-9a92aecfcc175d81fbba1a2961cf20fc01871424daf14c79863a98077305e85c): complete subsection reference.

<a id="canonical-9981f439ea97bf8cfebbf2fdd96afcedf5e2e4a0e312d39c65fb74380db72214"></a>

## Next pages — rule_list.rules.spec.segment_policy.src_segments / 373dadff407e / 4

- [rule_list.rules.spec.segment_policy.src_segments.segments](resources--service_policy--reference--group-003.md#canonical-9a92aecfcc175d81fbba1a2961cf20fc01871424daf14c79863a98077305e85c)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-9a92aecfcc175d81fbba1a2961cf20fc01871424daf14c79863a98077305e85c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eadb1168ca1131373995b13cf62cdddb213abab4969fe199ef7c1f96e1014c1f"></a>

## rule_list.rules.spec.segment_policy.src_segments.segments — rule_list.rules.spec.segment_policy.src_segments.segments / 5c88af7053e2 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b)
- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-7954f333c9f5786892f536436ac0d0c7d075d51971afe9d04e72e8684eaa5a4d)
- rule_list.rules.spec.segment_policy.src_segments.segments

<a id="canonical-122d5a0b00c86ec65928b724732b923f9c9ed424719e415da1201802c3ca29ee"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-52e54759b98d1e90cddbc7ca3fc719b8d93bef8c171bf69b5ee0dd85993133c3"></a>

## Direct properties — rule_list.rules.spec.segment_policy.src_segments.segments / 5c88af7053e2 / 3

<a id="canonical-3e75dc27315fbdd2b8cf1c32f50ac89a18351f239964afed81602787a86826af"></a>

<a id="canonical-396555dcbe8c83525fd0ffef3109df4f06e21bed7af1db82cd7ad3aa36c95a21"></a>

## name property — rule_list.rules.spec.segment_policy.src_segments.segments / 5c88af7053e2 / 4

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

<a id="canonical-2821f5aa43d1069604e3f1a526e14015e37d133585cd3c5cff4228c56f1c7939"></a>

<a id="canonical-f83b343586dabff56f3ca9d7cc6171c812effc42740029fc69e3a34588276475"></a>

## namespace property — rule_list.rules.spec.segment_policy.src_segments.segments / 5c88af7053e2 / 5

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

<a id="canonical-565593c228411b1cf6f7352af2481d86aac3bcbf708fecde9625b66026475613"></a>

<a id="canonical-5dab502464bd600fe240d943e3d8523fdfcc0b05e2dc97880ed8ad37eadc6797"></a>

## tenant property — rule_list.rules.spec.segment_policy.src_segments.segments / 5c88af7053e2 / 6

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

<a id="canonical-e834fef8c9419fe8369546f26057e17d4b6ec4e2acf08bbe0b36751ae88ac87f"></a>

## Next pages — rule_list.rules.spec.segment_policy.src_segments.segments / 5c88af7053e2 / 7

- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-7954f333c9f5786892f536436ac0d0c7d075d51971afe9d04e72e8684eaa5a4d)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-48e21c0002eed854bf2da383f1e5b04e92d13a8e1f01fc1c1b81d84a47db63e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34a2f43b8d088ed03158c299937e9849b232dd8f79e49508de137e3886d9daa5"></a>

## rule_list.rules.spec.tls_fingerprint_matcher — rule_list.rules.spec.tls_fingerprint_matcher / 48636e81e66f / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-a5cab8a51fdbe6dac16158e981f636009a68f81723a8a5c7f48ee47d9170c802"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-187d642e769d3b0354567d7d543a2f2edd8b1c833b2683957b9360a361d4b121"></a>

## Direct properties — rule_list.rules.spec.tls_fingerprint_matcher / 48636e81e66f / 3

<a id="canonical-3da192ed095934700ddea259e4740fb6bd9b984f39fb1109618f4c3733310ef7"></a>

<a id="canonical-3fc359db2d5cfa2310bb5b0c2f6b718f26f8369b7f03f775666b905023781014"></a>

## classes property — rule_list.rules.spec.tls_fingerprint_matcher / 48636e81e66f / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-780a1405e3da85d3b670cfba320783baf51f814c7cd567b07b626417016203a2"></a>

<a id="canonical-a4f7d514bc7495da5b575cc5d67d11968a9d0db0a1ddf2368fc7f0e0c948b965"></a>

## exact_values property — rule_list.rules.spec.tls_fingerprint_matcher / 48636e81e66f / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1ddbbc0154154787d19219a92604bef4c5ce92bc68c6050f8386e89604079b8b"></a>

<a id="canonical-b5fc80392816d396a821d42c0d37fe9e556a644e0bafce38819d191553f3cc81"></a>

## excluded_values property — rule_list.rules.spec.tls_fingerprint_matcher / 48636e81e66f / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-920fc3aa07e0dab8e5dde607b51dc1a87aedcc5c81a54bb422e65d7cbbf89f19"></a>

## Next pages — rule_list.rules.spec.tls_fingerprint_matcher / 48636e81e66f / 7

- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-c10fb3e61bf6eba8dd8ec88685f42e0566dcdd35111595de92aba354ff270132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90bc85c2b8002d9b25cc9336b2e442c63caa55c4e75cce78f6947c67b29cb04b"></a>

## rule_list.rules.spec.user_identity_matcher — rule_list.rules.spec.user_identity_matcher / 238518061459 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- rule_list.rules.spec.user_identity_matcher

<a id="canonical-6e129b1019240134a7ee3e9f32c1f7668f4dbd5806f53a16566b249705708e32"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
user_identity_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-5cdbc3e3ccd2ae6f5d6ad288ad1c8c5bce8b69d80a2c36402a8b9e06d0278f33"></a>

## Direct properties — rule_list.rules.spec.user_identity_matcher / 238518061459 / 3

<a id="canonical-79d03fa30b7d1a99f62e9577db3f948bd6a2dcab4d864abcc3fd2b09050537f1"></a>

<a id="canonical-f0b44b80aa3479276bcf9f801863437380b9620aad524fa0041607a4c1ba1961"></a>

## exact_values property — rule_list.rules.spec.user_identity_matcher / 238518061459 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8fde96738ceb16338cd03ec8241aa291c2b52b46bb23d94d8dca14b3296573ec"></a>

<a id="canonical-52d001431200061613cd2543656ffb43e7a47e1c2227355700d2ac8bad43a548"></a>

## regex_values property — rule_list.rules.spec.user_identity_matcher / 238518061459 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6c13327bb0e02cd8878221255cb05bb7b4faa2aaf3c1128c082780187a153ac6"></a>

## Next pages — rule_list.rules.spec.user_identity_matcher / 238518061459 / 6

- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f96e140102fb5355579b425449de025cd6fad96d8ce06c94ea415e3dcdc667f7"></a>

## rule_list.rules.spec.waf_action — rule_list.rules.spec.waf_action / f689ada9d93b / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- rule_list.rules.spec.waf_action

<a id="canonical-bb74d13ec22c2e1cbaaa5036dc9d05c5b80964f1913f42605c5cf7232fa7fb8b"></a>

Type: `"object"`. single nested block, Optional.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "none"),
  validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingObjectAttributes("none",
    "waf_skip_processing")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

Terraform syntax:

```terraform
waf_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-e5c6d14fa2ab07f47adaa8ebbbc5bb44d194f00f3d140dec000c87017b03d1d4"></a>

## Direct properties — rule_list.rules.spec.waf_action / f689ada9d93b / 3

- [app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5): complete subsection reference.

- [none](resources--service_policy--reference--group-003.md#canonical-233a327091ac56a08510880d590e18b8d97d2d9b2ca4cbbb3d038b89fcd5d768): complete subsection reference.

- [waf_skip_processing](resources--service_policy--reference--group-003.md#canonical-3ec28ee77beccaba5dd80f2aae381531269dda2239717caf904abbe78e1c162c): complete subsection reference.

<a id="canonical-64739c1e9052dc752df8bff5988d112911188d73e67590512402210124873a0f"></a>

## Next pages — rule_list.rules.spec.waf_action / f689ada9d93b / 4

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- [rule_list.rules.spec.waf_action.none](resources--service_policy--reference--group-003.md#canonical-233a327091ac56a08510880d590e18b8d97d2d9b2ca4cbbb3d038b89fcd5d768)
- [rule_list.rules.spec.waf_action.waf_skip_processing](resources--service_policy--reference--group-003.md#canonical-3ec28ee77beccaba5dd80f2aae381531269dda2239717caf904abbe78e1c162c)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47104aece4d436c960d2eb51f203f680db29e37e76ca91de7b6312158b5a12da"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control — rule_list.rules.spec.waf_action.app_firewall_detection_control / e38d5d4d59ee / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- rule_list.rules.spec.waf_action.app_firewall_detection_control

<a id="canonical-ece11ac74536b3ce1169bc10d08dc22abef3a0dfcf85ecf5a60da52b972ec400"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-f6c1e35cf81defcd64ad5b61d58e62fa07eb3d2835283398ce9001a13fa11aba"></a>

## Direct properties — rule_list.rules.spec.waf_action.app_firewall_detection_control / e38d5d4d59ee / 3

- [exclude_attack_type_contexts](resources--service_policy--reference--group-003.md#canonical-5982191ed6225356c8fea7a232ec7e57c6b15fe30f58a856ec1bdc9266c07525): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy--reference--group-003.md#canonical-15d07c7260e07a5069df08dd0dbf5e14b5031c1a4be5ee2e66028854db47d93c): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy--reference--group-003.md#canonical-acb8c582df9e217aad819c963e5691c335a105da70f0d097e7b817c8ee5baa25): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy--reference--group-003.md#canonical-ddda747b9e78f83a56d2892194884192eeddc8f19aa7fdc161f934b588b31253): complete subsection reference.

<a id="canonical-d134440ed20fba6d60bc502ae0f940f597a721d7e337b922bee029bfce02a776"></a>

## Next pages — rule_list.rules.spec.waf_action.app_firewall_detection_control / e38d5d4d59ee / 4

- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy--reference--group-003.md#canonical-5982191ed6225356c8fea7a232ec7e57c6b15fe30f58a856ec1bdc9266c07525)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy--reference--group-003.md#canonical-15d07c7260e07a5069df08dd0dbf5e14b5031c1a4be5ee2e66028854db47d93c)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy--reference--group-003.md#canonical-acb8c582df9e217aad819c963e5691c335a105da70f0d097e7b817c8ee5baa25)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy--reference--group-003.md#canonical-ddda747b9e78f83a56d2892194884192eeddc8f19aa7fdc161f934b588b31253)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-5982191ed6225356c8fea7a232ec7e57c6b15fe30f58a856ec1bdc9266c07525"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b039304ac3cb59e959c306a393dbb5ed5382bb09b0b8fcd1787d8a239a590b7"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_ty / 5ba820658828 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-e3546643c2eedb4dac4e07645701fc5fbf0a80d278b03ef02724ff700a5b6610"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-17f1488f3e994ba5302026668b0d189dab45537e158240743a9af9d25708b0ab"></a>

## Direct properties — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_ty / 5ba820658828 / 3

<a id="canonical-af9744104cea10643baeb9da4aebed1817021eec78ade07f2f8bd00763a3e3b0"></a>

<a id="canonical-b7c3feda3d019f18d35c7fb7024b3b9325f54b92ba9d2b0717b0092981a9fe4f"></a>

## context property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_ty / 5ba820658828 / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-71f4540d9ec543c0f7630b5950f318490830a8c88b4fc6af91ea5f874bbf8ba0"></a>

<a id="canonical-3a233566c95b1ab7da5b4d998f2c4ef926efa8017bfb6ff17bb07dc0fd33b4b0"></a>

## context_name property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_ty / 5ba820658828 / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-25b76d441c1d596286c3bd457a9faf3c6456acf6d63fc7f28a3e30f9a7d319d3"></a>

<a id="canonical-c82540b993549b9ac41de4875c2d7deb15a92aab719a25c999b7e4700f81f40b"></a>

## exclude_attack_type property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_ty / 5ba820658828 / 6

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cd5eb2b0df95ac0968860b61e4bf7c702933d4dc7a7ddcf2d109a65d4318afb7"></a>

## Next pages — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_ty / 5ba820658828 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-15d07c7260e07a5069df08dd0dbf5e14b5031c1a4be5ee2e66028854db47d93c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0c31b20b66da7034103bf0b19b7450c9df38a60f5127e422162611f65729ee6"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_ / 071a2334ee81 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-f05567d934380ea1c7f54218b296280578e4689d5b4d15d2bd8004b092eb921c"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c2659ed1dcfb3c574adc248612268728ea512988c6617492fc2229aa3265a56"></a>

## Direct properties — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_ / 071a2334ee81 / 3

<a id="canonical-807636ee3826b9aedc7e7903f75c30cf584c00451bc08b7429b110abb60b7543"></a>

<a id="canonical-110f8a0eecb0179d2cdb6cbc762b9ae6dd7d328d85a06cfdbbe88a10ed7e477c"></a>

## bot_name property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_ / 071a2334ee81 / 4

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-3a4c079d1a3b6828c853a71ce6e0721ac0ad4e8672d074932e70671c19e39cb6"></a>

## Next pages — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_ / 071a2334ee81 / 5

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-acb8c582df9e217aad819c963e5691c335a105da70f0d097e7b817c8ee5baa25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-052b15cee66243951804ce1af57291a7889b887789e7db8d615ed06b84557cda"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature / 1a4f819106dc / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-47006d64ec1b0e838dae87121fe7fd29918cb94b8288527ba3d85097c35825ef"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f16a0a3066fd8b352ad15f7a5b0d2a07f4a908db639f010266d98af707ffc6cb"></a>

## Direct properties — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature / 1a4f819106dc / 3

<a id="canonical-4006bb4c5e4721f8759121e5d85e2913bcfad025afb31b857333e1ed14e5c117"></a>

<a id="canonical-8f5fcef33bd13b5bacfd57e8c7ef5175a1d3634854e93d708b86e5311295f709"></a>

## context property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature / 1a4f819106dc / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e3984328ac6dc620d444a928e540cf071c7d936959d87366aefc74f2f82d2af3"></a>

<a id="canonical-d6ac6f31ad563ffc983a6f01f40c11f4a182189999b97a68187b461543e817db"></a>

## context_name property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature / 1a4f819106dc / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-7ca8c19e533362228fd0de50311dd05ff2e35ebe5cd490fa9543b7f2b9e95c89"></a>

<a id="canonical-d6b0976e2848ea877329ce38b7b94957ba2bfd1b92d2a82829727e9c60f012e6"></a>

## signature_id property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature / 1a4f819106dc / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-427bcc6ad96d82d612bd3bb376f374e76e0139aa90e42772765a582b73b3b045"></a>

## Next pages — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature / 1a4f819106dc / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-ddda747b9e78f83a56d2892194884192eeddc8f19aa7fdc161f934b588b31253"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2000b29d49edac221e6e35e41bc1046077af96e42d7c918fbb5b2c5e20bfcdc"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation / 7b1601f03c8a / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-50a7ae6da673301ca3fec365e62936b2789b8011d8eaec9de856eddca797d74e"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-5f05a36da624d281421bcc5b6ccbe7f1a9501cf6b6445da7ac9b2fa3f24291ff"></a>

## Direct properties — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation / 7b1601f03c8a / 3

<a id="canonical-71bbf0c7365b6cf9565039f83cb510707f48c910dd672462816776fbe64754c0"></a>

<a id="canonical-7af30f0255a90823a279721b7f6b60ddf5575c3ecbcbd7a338b82980059fb573"></a>

## context property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation / 7b1601f03c8a / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-730ecb6406476ae187160830005eb1b0ff1ed536b1625de5550c5be3bd180e50"></a>

<a id="canonical-ea863096c1b58b008fd0a740029c6b287b478d83d68a5a026bd0d12fbf9a91d9"></a>

## context_name property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation / 7b1601f03c8a / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-f7d72fd9effab4fe20eeb60febcfc54d3ffa651dc93abaf0fa763c43d626b13b"></a>

<a id="canonical-6b39669418741e4429d988b2bfa4222b2a3e220836ef8f95619a7b0500f1a86a"></a>

## exclude_violation property — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation / 7b1601f03c8a / 6

Type: `"string"`. Optional.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7aeefcf74a69e5d0f2e004a87710f7cb83f0d5f93e34cbf52932da6fc1f4abb7"></a>

## Next pages — rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation / 7b1601f03c8a / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-cae791e1083f268eb5171c3b6251e50a6fc69fd60c62c43ad3ed54c1641f1bc5)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-233a327091ac56a08510880d590e18b8d97d2d9b2ca4cbbb3d038b89fcd5d768"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0ec6ea470b950e6574aa21771126aa69a58474f043aa6bbbc17ea9b193431f0"></a>

## rule_list.rules.spec.waf_action.none — rule_list.rules.spec.waf_action.none / 0f0db0531820 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- rule_list.rules.spec.waf_action.none

<a id="canonical-5c95e10f062fff1bdadaafe2b301e3df8c1ac9441bc29b6b474dc17bfe451018"></a>

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
none = {}
```

<a id="canonical-52556185d960506cb56ec24237ba42402c27018c7f26c09268d68252fa7b7e57"></a>

## Direct properties — rule_list.rules.spec.waf_action.none / 0f0db0531820 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a3a2f73fe4ddfbd1f086266bb855c0b2f8edd625ae3f731b783bb3f5ee4330c"></a>

## Next pages — rule_list.rules.spec.waf_action.none / 0f0db0531820 / 4

- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-3ec28ee77beccaba5dd80f2aae381531269dda2239717caf904abbe78e1c162c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f74e90bbd7e077a23a04dcee91a9c8eefa71e9a79522fdaf16159353215be9c5"></a>

## rule_list.rules.spec.waf_action.waf_skip_processing — rule_list.rules.spec.waf_action.waf_skip_processing / fb85ff89dac7 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- rule_list.rules.spec.waf_action.waf_skip_processing

<a id="canonical-3ef89b84f10229bef7f6f40e1e9819a0fdc3ea058e0753739874f9fe6ada3f8a"></a>

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
waf_skip_processing = {}
```

<a id="canonical-59039986444c7ea83cbee9d9d16a0f1df280d85a9401c7aae85e5837370bab7a"></a>

## Direct properties — rule_list.rules.spec.waf_action.waf_skip_processing / fb85ff89dac7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-430e3ac78146b593f7d9515b84bdf79e61661f62c32184305e7791b12993b6e8"></a>

## Next pages — rule_list.rules.spec.waf_action.waf_skip_processing / fb85ff89dac7 / 4

- [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-de3685065567f4411ab04f86f8f14f9fc4e21e79652d75a6706a62e3882309c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-559cfe159be7a78ba8288830a0e921a125d0c0f4c8c3c7130238ab86373d8598"></a>

## server_name_matcher — server_name_matcher / fafc1240de11 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- server_name_matcher

<a id="canonical-80e3e5fd99369410697e0682fa18cd97d32af4c850aa895248f5a56b23e38066"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
server_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3caca898c31124c761cf4dc4b25d9e75fd327c4f65cc85751302b7baf3471eed"></a>

## Direct properties — server_name_matcher / fafc1240de11 / 3

<a id="canonical-3e049eb59da2f103404aef24cb640e4e5a14c4cae92ddf732af6d33fa05aebed"></a>

<a id="canonical-bbd0f81ee86492142b4283bd79c473000c7a32e56c26d09fe3a8a44f2f05de77"></a>

## exact_values property — server_name_matcher / fafc1240de11 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5c3098b14d6a661e9dc45ea8f16df74f7d1061f5b85f23ba6d61bb8a67488f04"></a>

<a id="canonical-afdad96a8dde12c5299c9591289cfe7900cebc484b3807673121f4725d3a30b6"></a>

## regex_values property — server_name_matcher / fafc1240de11 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b27260359c92aa4aaf192a48b853fce295516b42d4503a1ae474981288a2ebb1"></a>

## Next pages — server_name_matcher / fafc1240de11 / 6

- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-ce09b78b2268dfdf225f8fa7808ee3dc2ee2eaa1ca7f4aef27737f6280575fab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ea499145db6919b04138ac75778561232d993a8092d49b9088de830f928e827"></a>

## server_selector — server_selector / c2d0b286408c / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- server_selector

<a id="canonical-904ee2af17440d71cdf493a1ce98c30af916155ceb30517d84c371e051cba909"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
server_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-29096758e85777811096091585622794cdf00cc0c1e6503603608a3c59afe3ea"></a>

## Direct properties — server_selector / c2d0b286408c / 3

<a id="canonical-2b0f1323ad1451b6ecc9179323a2631e1fc220e24f536a6c057573242605d3e8"></a>

<a id="canonical-548a0d757891b4ee18ed787ac806d34b4697e8c4459638f21555c322e3517b5a"></a>

## expressions property — server_selector / c2d0b286408c / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-12c840270ba90a28f841d73474316e9f7c7c6ad790a4703f06b35b9f0c305544"></a>

## Next pages — server_selector / c2d0b286408c / 5

- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-4c2b79645b86fc0db1b5a5e8348eb1539ad05a418b46eed14d7fd13168c502a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5cb7f8f850fb93c78e733b52d1012a5af1634288f9e50c1add187c88d4f8037"></a>

## timeouts — timeouts / d5d3b1c7677c / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- timeouts

<a id="canonical-bef30de4cea2cb42a40409f558317d61fb10e72631ebed266bc38d1d4292c923"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5128100870d599af2d6ffdac17a5b58537ecda26b0f4999b163ffd2fb234993"></a>

## Direct properties — timeouts / d5d3b1c7677c / 3

<a id="canonical-1caffda87097fdedc967c3a9b2b835c5e3be69e720682e2ba705339db493200c"></a>

<a id="canonical-d7b96e3b78f8991d1fdc0fe146d660867a7a310a5de4cb3d51d04840e45243f1"></a>

## create property — timeouts / d5d3b1c7677c / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4dc9b459511c80df0eda975950cc73547a7c4aa8c1f31caa73c0f147e50e59aa"></a>

<a id="canonical-fa44bd4ca74c35434b1a6d25056fb818a9d2717216a5a749738d7dfb57a1b6cd"></a>

## delete property — timeouts / d5d3b1c7677c / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-eafb5d0d70c439a9599f9c6db37cf30f1d62bda11c5c7712ab651497c4caa7e6"></a>

<a id="canonical-f197bd471795762aa5708044581fea90cfdaeb6d01bd804ef9b6394235360a44"></a>

## read property — timeouts / d5d3b1c7677c / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-ab3e7ee70b5f533bacba83b7261465348be4231e2e06e8801981a097b00771fa"></a>

<a id="canonical-b0d7bf5dc088e891877d828fbc81fe6030c8a0027b58991d32250a1b07250c44"></a>

## update property — timeouts / d5d3b1c7677c / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-eb54c90a4a1abaec30a64dbc24d6ced2eb437075e684fef2fe575cc0ba636175"></a>

## Next pages — timeouts / d5d3b1c7677c / 8

- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
