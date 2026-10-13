from pathlib import Path
import json,os
patterns={
 'virtual_site_required':['virtual site','virtual_site'],
 'bfd_required':['bfd'],
 're_fallback_required':['fallback'],
 'control_plane_required':['control'],
 'data_plane_required':['data plane','data_plane'],
 'site_selection_required':['site'],
 'name_limit':['length','name length','too long'],
 'mesh_type_required':['mesh','type'],
 'missing_required_input':['required','must','cannot','empty','missing'],
 'unsupported':['not supported','unsupported'],
 'reference_missing':['not found','does not exist'],
 'namespace_constraint':['namespace'],
}
raw=json.loads((Path(os.environ['XCSH_ACC_PRIVATE'])/'mesh-probe.json').read_text())
rows=[]
for row in raw:
 variant=row['variant']
 if variant not in ['base','markers','data_plane']:raise ValueError('unexpected mesh variant')
 error=row.get('error','').lower()
 rows.append({'variant':variant,'accepted':row['accepted'],'signatures':[key for key,terms in patterns.items() if any(term in error for term in terms)]})
Path('test-reports/mesh-probe.json').write_text(json.dumps({'version':1,'variants':rows},indent=2)+'\n')
