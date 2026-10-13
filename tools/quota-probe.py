import json,os,urllib.request
from pathlib import Path
targets=['api_crawler','crl','container_registry','global_log_receiver','tcp_loadbalancer','udp_loadbalancer','forwarding_class']
rows=[]
for namespace in ['system','shared']:
 request=urllib.request.Request(os.environ['XCSH_API_URL'].rstrip('/')+'/api/web/namespaces/'+namespace+'/quota/limits',headers={'Authorization':'APIToken '+os.environ['XCSH_API_TOKEN'],'Accept':'application/json'})
 with urllib.request.urlopen(request,timeout=30) as response:document=json.load(response)
 Path(os.environ['XCSH_ACC_PRIVATE']+'/quota-'+namespace+'.json').write_text(json.dumps(document))
 for target in targets:
  row={'namespace':namespace,'resource_type':target,'matching_entries':[]}
  for section in ['objects','quota_usage','resources','float_quota_usage']:
   group=document.get(section)
   if not isinstance(group,dict):continue
   for key,value in group.items():
    if key not in [target,'ves.io.schema.'+target,'ves.io.schema.views.'+target]:continue
    item={'section':section,'value_type':type(value).__name__}
    if isinstance(value,(int,float)):item['value']=value
    if isinstance(value,dict):
     item['field_count']=len(value)
     item['numbers']={name:number for name,number in value.items() if name in ['limit','usage','used','max','maximum','value','current','available','quota','count','remaining'] and isinstance(number,(int,float))}
     item['text_shapes']={name:type(field).__name__ for name,field in value.items() if name in ['limit','usage','used','max','maximum','value','current','available','quota','count','remaining','name']}
    row['matching_entries'].append(item)
  rows.append(row)
Path('test-reports/quota-evidence.json').write_text(json.dumps({'version':1,'resources':rows},indent=2)+'\n')
