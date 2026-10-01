---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-3223311211232323-1131320131230232-3000301210020221-0330320132202223-3301003331223331-2131131233021211-1001132202033323-2121020023112030"></a>

## voltstack_cluster.site_local_network.new_network — new_network / 200003330021 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- voltstack_cluster.site_local_network.new_network

<a id="canonical-0203112011300100-2232221113202223-1002330210332203-2323131202103001-2221321201022013-1200123201221112-3201302030013331-0310330102320111"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233120333011103-3200230022330312-0231023320000310-0203302221202310-0200331010203212-3133123233321210-1020331333020031-0002212001212103"></a>

## Direct properties — new_network / 200003330021 / 3

<a id="canonical-0331302222121311-0303133100223212-0330030202332101-3202013320130222-2330020123013102-0032323300321320-1212122330231330-3030131231211123"></a>

<a id="canonical-0200213101320121-1223211033220302-3030211303221203-2320222120120213-2210020131010120-2122030013002212-2021111023110010-3230233233223322"></a>

## name property — new_network / 200003330021 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3310030001012101-1120123011203031-1210323130330032-1333132212203222-0310122030323233-3302121021320020-2100300130200221-0133320313110212"></a>

## Next pages — new_network / 200003330021 / 5

- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2221021031302100-2310231033112111-2333333110213200-3212311123133201-2011221231020333-3123122102322211-1000330301223202-0320111111311131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202022031331032-2201000112211013-0220030212122311-2301330312123232-0212212201110301-1123103302213330-3031022213333230-3032330331231323"></a>

## voltstack_cluster.site_local_network.new_network_autogenerate — new_network_autogenerate / 311212310332 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- voltstack_cluster.site_local_network.new_network_autogenerate

<a id="canonical-0003300001200311-3103132200321101-3312300211211302-0010203021333111-1010132023120233-2001001303331310-3301300111322331-0331020313232300"></a>

Type: `["object", {}]`. Optional.

Create a new GCP VPC Network with autogenerated name.

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
new_network_autogenerate = {}
```

<a id="canonical-1212321321311020-2130313012311113-3232331332110121-3103012000301130-2232112311201201-0203133333211011-1303113221120023-2001301103031000"></a>

## Direct properties — new_network_autogenerate / 311212310332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331113201231322-1101022022110232-3300012030031112-2013020321000002-0123320322002033-3223120112303132-2213233010022023-2123013331222200"></a>

## Next pages — new_network_autogenerate / 311212310332 / 4

- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012203131211301-0300203131121211-3002011032111312-3223203333220130-3312230320012211-3321222332002020-1212003200310201-0332032203313101"></a>

## voltstack_cluster.site_local_subnet — site_local_subnet / 201221332012 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.site_local_subnet

<a id="canonical-1112233310222231-3323120000012103-3331013012230112-2010330332111120-1332030031220222-3312133112303103-1200201133331310-2102020021121331"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
site_local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310002213021331-3333102113321002-0111321023212321-3232322321112011-3001101301311232-0003133213210110-2313310110011102-1101130311002021"></a>

## Direct properties — site_local_subnet / 201221332012 / 3

- [existing_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1311312302212032-1030302113103213-0233300303132201-1303222322122032-1223223023300323-0000333210120000-0321332100223310-3232332111102333): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-3020211232002203-2010302022121123-2011333211121332-1001233012200102-3230312231031213-2103323301100000-3112113330333223-0322123222332310): complete subsection reference.

<a id="canonical-2302302012322330-3210030302011300-3132100111211311-2110020003233302-3031232030331222-3023221231101000-3101201233322023-2030201130301003"></a>

## Next pages — site_local_subnet / 201221332012 / 4

- [voltstack_cluster.site_local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1311312302212032-1030302113103213-0233300303132201-1303222322122032-1223223023300323-0000333210120000-0321332100223310-3232332111102333)
- [voltstack_cluster.site_local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-3020211232002203-2010302022121123-2011333211121332-1001233012200102-3230312231031213-2103323301100000-3112113330333223-0322123222332310)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1311312302212032-1030302113103213-0233300303132201-1303222322122032-1223223023300323-0000333210120000-0321332100223310-3232332111102333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002033021220122-3002312232211022-1303212010102301-2122303121312321-1302011101200313-3113003032112310-0203202100301202-3230332121311333"></a>

## voltstack_cluster.site_local_subnet.existing_subnet — existing_subnet / 022332223311 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233)
- voltstack_cluster.site_local_subnet.existing_subnet

<a id="canonical-0010213320213230-1200122022123022-0013100100123331-1102011121101112-3301130302333201-0023203120131302-2103221301101013-1011330133131230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203230022103223-2110231331133022-0032131321000021-0333110203022231-1233021221310112-3030301311022030-0221323133003310-3232221013213132"></a>

## Direct properties — existing_subnet / 022332223311 / 3

<a id="canonical-2110213000321012-3322030300311112-0223230113003003-3123232330312010-0232312320221233-1032321311333231-1211032303003121-1003033121300123"></a>

<a id="canonical-3010020013133233-1002232202303203-2110221210301131-0111113101133300-0021000331222200-1223021130111210-2313311120231122-3103020021123323"></a>

## subnet_name property — existing_subnet / 022332223311 / 4

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0003331021232301-2133112300031122-1032001111103102-2011011133113211-3020010300003011-2110102023033302-0221231013013021-1230000020123130"></a>

## Next pages — existing_subnet / 022332223311 / 5

- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3020211232002203-2010302022121123-2011333211121332-1001233012200102-3230312231031213-2103323301100000-3112113330333223-0322123222332310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303321201222323-3231021113200320-1233331130312221-0223001333321102-2203230332123122-3302113301030100-0313310100111003-0133331022231322"></a>

## voltstack_cluster.site_local_subnet.new_subnet — new_subnet / 033123133012 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233)
- voltstack_cluster.site_local_subnet.new_subnet

<a id="canonical-2333203110321011-3221113220130113-3223003100020201-2030130112012202-3212130130113220-1201111131012022-2112302200310330-2230313121002000"></a>

Type: `"object"`. single nested block, Optional.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4")}
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
new_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233033323222321-1323030012233021-3103031131220332-0233312011311323-0300030333121232-3130223300330100-2103311123302233-0013021232223003"></a>

## Direct properties — new_subnet / 033123133012 / 3

<a id="canonical-0121313112032323-1030012112032203-3211102210303310-2033332323103111-3020203113012112-0122230200212300-1312001200303120-3212002132311002"></a>

<a id="canonical-2022011022100100-3011200222021021-2233212332113010-3212123010210021-3233223102230313-0223010010310020-3133230333222001-1130221313123012"></a>

## primary_ipv4 property — new_subnet / 033123133012 / 4

Type: `"string"`. Optional.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-1022333112000021-3223132022301030-2213100133322122-2032121110002003-0323221310013213-0110101001132002-1322232120030212-2301022300123012"></a>

<a id="canonical-2103102120312003-3321213210113333-0121033231112211-1232020223023210-3333300200032213-1220102311122212-2313321331213230-1103211120230111"></a>

## subnet_name property — new_subnet / 033123133012 / 5

Type: `"string"`. Optional.

Name of new VPC Subnet, will be autogenerated if empty.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3122213022113301-3301012232112122-2001011311003201-3230321231030111-3011012300210010-2233022221201221-3200102231211303-2300211222231212"></a>

## Next pages — new_subnet / 033123133012 / 6

- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2110313330032303-3110333222121201-2220303222032313-0202012020100132-0013113323310002-2221121220013110-3121003203331012-0010232203321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323330023112203-2011101101013110-3133032303102103-1301331211031321-0021011322323233-2233022202301331-3132112301201120-0330310113110300"></a>

## voltstack_cluster.sm_connection_public_ip — sm_connection_public_ip / 232130300321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-1330023103101131-1323220103232030-3323301010002103-2210330220321133-2110031213333013-1010031200101022-3202033111010302-3203311032110223"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-2023020332220121-2210131210130231-0301101133002232-1220031012300012-0100323320321313-1133130320303330-0231303302033133-2000100231031002"></a>

## Direct properties — sm_connection_public_ip / 232130300321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110310023132113-3021212201221123-2210233330333303-0121103331221211-1210221001200002-1310111010103013-2033203121303302-3220113203030103"></a>

## Next pages — sm_connection_public_ip / 232130300321 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0331101100010033-0222011110310013-2032302133123312-0220133021132300-2101231320333020-0020000111202113-1303322132113030-2132013223233320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000000120232202-2230022122021332-0320010310100031-3300013322231203-3003130110122111-1112211003203032-2011320021031110-2323133021123220"></a>

## voltstack_cluster.sm_connection_pvt_ip — sm_connection_pvt_ip / 210220201113 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-0001231201231133-1001001202203233-3311113131230131-1221031000210331-2301230312233001-3102033110232132-0331011113322022-1211011311231002"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-1313010221100300-0203021021222001-0330232211102223-2110310000320231-3013011131211100-3112103032232003-3103311031023221-0312001022312213"></a>

## Direct properties — sm_connection_pvt_ip / 210220201113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123222210331121-1233301310233130-3113301230310321-3222023130123323-0200302001031103-0010123132310321-3003110331311230-1313121010022232"></a>

## Next pages — sm_connection_pvt_ip / 210220201113 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1112333322110321-0200101320023311-3310121222132330-0033000133132300-1232002332132331-2111300311231333-1211022132123111-1313233202133121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031301023203100-2022110003311210-0223220300231100-1000301200200210-2221031320021112-1020132110001322-3222320133112233-3211322313131032"></a>

## voltstack_cluster.storage_class_list — storage_class_list / 012300112100 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.storage_class_list

<a id="canonical-2012112022110310-3212303113133310-2000330201133312-1111100233121020-2021201323101033-1001130210331301-2123122110012312-0013212312300321"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

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
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122230200112300-3210203302220102-3032013303333221-0121103313202333-1033132321220130-2003220320223021-1321010312032013-2001220102120032"></a>

## Direct properties — storage_class_list / 012300112100 / 3

- [storage_classes](resources--gcp_vpc_site--reference--group-005.md#canonical-3033200112002300-0313122330223123-3101332322122332-3132022123203213-3220101230013330-1233322113303010-1100221002032000-1223012001322112): complete subsection reference.

<a id="canonical-2232111030201110-1033220323000003-3231031102321203-1221220031012021-1301303022233210-0112302022112100-2232003321022110-3300110002132123"></a>

## Next pages — storage_class_list / 012300112100 / 4

- [voltstack_cluster.storage_class_list.storage_classes](resources--gcp_vpc_site--reference--group-005.md#canonical-3033200112002300-0313122330223123-3101332322122332-3132022123203213-3220101230013330-1233322113303010-1100221002032000-1223012001322112)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3033200112002300-0313122330223123-3101332322122332-3132022123203213-3220101230013330-1233322113303010-1100221002032000-1223012001322112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013011111203300-0220232110322030-1330223210211313-0020122300022113-0302223322230122-1002120301301221-0012302211211123-2233223301013232"></a>

## voltstack_cluster.storage_class_list.storage_classes — storage_classes / 233221123221 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-005.md#canonical-1112333322110321-0200101320023311-3310121222132330-0033000133132300-1232002332132331-2111300311231333-1211022132123111-1313233202133121)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-1233311323031330-3111022113232202-0221021202303321-2210300122211231-0030023331022333-3131120113202002-3122001201330231-1120120010112003"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021321330122212-3033220020203113-0011121101230021-3120222230322201-3013121333100033-1032231120110330-3122302303212331-0131211120131320"></a>

## Direct properties — storage_classes / 233221123221 / 3

<a id="canonical-1013230323222321-1201113220332032-3000020121233231-1333211000101323-1100333310201212-0321311231222121-3123100011120320-2231331212213003"></a>

<a id="canonical-2131230123210333-2031121302123002-3111332010331003-1020202200321333-1333312300022230-1103210000022020-0312320111111220-2222310321313311"></a>

## default_storage_class property — storage_classes / 233221123221 / 4

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

<a id="canonical-2031121313222202-3123233001232312-3223303212130300-0212223230011330-1003132330201221-3202113320222202-3022320302123020-3213221113033210"></a>

<a id="canonical-0211221112330021-1311103330110232-2101210130013221-0102330230223230-3021110310332201-0113120231001300-1300131303201122-3101232112120023"></a>

## storage_class_name property — storage_classes / 233221123221 / 5

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

<a id="canonical-0201230101223002-2112332330311230-1320203213123002-3332110110001302-1100301232010320-0332113320200113-3002231021331113-0103122212001121"></a>

## Next pages — storage_classes / 233221123221 / 6

- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-005.md#canonical-1112333322110321-0200101320023311-3310121222132330-0033000133132300-1232002332132331-2111300311231333-1211022132123111-1313233202133121)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022221002321233-3301311022331203-2023200300013320-3100210310312032-3221002130120223-0030021101131103-1133202032012012-0102322223113322"></a>

## waf_signatures — waf_signatures / 121011231210 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- waf_signatures

<a id="canonical-0200303132233121-2101103333000100-0121030102330013-2121200130113333-1302301003223332-2320122122320002-2210031023012000-1021113133221212"></a>

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

<a id="canonical-0032323200000120-2333022203130202-2202123303111102-1310102000232013-1033203313133100-0022311131000122-0213211011230022-1223102210131002"></a>

## Direct properties — waf_signatures / 121011231210 / 3

- [automatic](resources--gcp_vpc_site--reference--group-005.md#canonical-1313300332013133-0303111131323211-1330301232330132-2230003021121032-1100231030021103-3120030323130321-3112331010320323-1122202110300120): complete subsection reference.

- [manual](resources--gcp_vpc_site--reference--group-005.md#canonical-1122000123011231-1321120132201013-0012201213120031-0302120102013020-3101112110112122-3331212022123203-2333112112110131-1010221221032022): complete subsection reference.

<a id="canonical-1010102231310222-0133303200313210-1031120211102103-0032303303123310-1323121130023211-3321231012303312-0310132021113332-3300333023211222"></a>

## Next pages — waf_signatures / 121011231210 / 4

- [waf_signatures.automatic](resources--gcp_vpc_site--reference--group-005.md#canonical-1313300332013133-0303111131323211-1330301232330132-2230003021121032-1100231030021103-3120030323130321-3112331010320323-1122202110300120)
- [waf_signatures.manual](resources--gcp_vpc_site--reference--group-005.md#canonical-1122000123011231-1321120132201013-0012201213120031-0302120102013020-3101112110112122-3331212022123203-2333112112110131-1010221221032022)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1313300332013133-0303111131323211-1330301232330132-2230003021121032-1100231030021103-3120030323130321-3112331010320323-1122202110300120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231233131112231-1110223011112312-0123032302003211-2032322010121233-0213220301210202-0230211023023032-2320020130333333-2312100011321101"></a>

## waf_signatures.automatic — automatic / 332320011312 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100)
- waf_signatures.automatic

<a id="canonical-1321103100021330-1221211333223021-1203223102233311-0102130330003133-2130331112110030-3131033032210320-2033332233131233-0003103320233211"></a>

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

<a id="canonical-2320303002310031-0332211100203001-2332013213102220-0200323030113120-3220332110130310-1301120212120132-2131230230312112-1033213203300100"></a>

## Direct properties — automatic / 332320011312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212230331212123-2311331112300101-2121222103103312-1112101111230321-3031131201330331-3021321301331203-0123002230012113-3310231320121103"></a>

## Next pages — automatic / 332320011312 / 4

- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1122000123011231-1321120132201013-0012201213120031-0302120102013020-3101112110112122-3331212022123203-2333112112110131-1010221221032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020101200301121-0203033232011112-2210222221030120-1230111211330012-0120000230133031-1322330123313232-3221301321112033-1013100333223020"></a>

## waf_signatures.manual — manual / 130321122300 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100)
- waf_signatures.manual

<a id="canonical-3301321103032221-0303320300031322-2111301110210302-1331000211302200-3323323133321311-3002113320123013-2322303002133100-1230010010130103"></a>

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

<a id="canonical-3121111132113013-1123023300200012-2313101032021230-1121120300210200-0022022323321232-0233220221131020-2131101231133332-2320022333130232"></a>

## Direct properties — manual / 130321122300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331111323313222-0121103323313230-0021312333000121-1303130110322010-2213121003320322-1303201200231223-0131300000003023-0120221111121002"></a>

## Next pages — manual / 130321122300 / 4

- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-2220233032031010-1323100300312032-0222201202010130-3130203022232222-0300133130321233-1132212212212321-0322012021310311-1132132110223100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
