select '<a href="/feat/'||alias||'" class="link">'||upper(substring(name, 1, 1))||lower(substring(name, 2))||'</a>' from feat f where lower(alias) = lower(replace('Improved Familiar', ' ', ''));
select '<a href="/skill/'||alias||'" class="link">'||upper(substring(name, 1, 1))||lower(substring(name, 2))||'</a>' from skill s where lower(alias) = lower(replace('Climb', ' ', ''));
select '<a href="/spell/'||alias||'" class="link">'||lower(name)||'</a>' from spell s where lower(alias) = lower(replace('magic missile', ' ', ''));
select '<a href="/magicItems/weapons?scroll='||alias||'" class="link">'||lower(name)||'</a>' from magic_item_ability mia where lower(alias) = lower(replace('corrosive', ' ', ''));
select '<a href="/bestiary/appendix/universalMonsterRules?scroll='||alias||'" class="link">'||lower(name)||'</a>' from monster_ability ma where lower(alias) = lower(replace('compression', ' ', ''));
select '<a href="/class/'||alias||'" class="link">'||lower(name)||'</a>' from "class" c where lower(alias) = lower(replace('Alchemist', ' ', ''));
select '<a href="/race/'||alias||'" class="link">'||lower(name)||'</a>' from race r where lower(alias) = lower(replace('human', ' ', ''));
select '<a href="/weapons?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias||'" class="link">'||lower(name)||'</a>' from weapon where lower(alias) = lower(replace('crossbowlight', ' ', ''));
select '<a href="/armors?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias||'" class="link">'||lower(name)||'</a>' from armor where lower(alias) = lower(replace('full plate', ' ', ''));
select '<a href="/bestiary/beast/'||alias||'" class="link">'||upper(substring(name, 1, 1))||lower(substring(name, 2))||'</a>' from beast b where lower(alias) = lower(replace('Aranea', ' ', ''));
select '<a href="/bestiary/appendix/creatureTypes?scroll='||alias||'" class="link">'||lower(name)||'</a>' from creature_type ct where lower(alias) = lower(replace('plant', ' ', ''));

for replace rules
select '"'||s.eng_name||' link": '||jsonb_build_object('find', s.eng_name, 'replace', '<a href="/skill/'||alias||'" class="link">'||upper(substring(name, 1, 1))||lower(substring(name, 2))||'</a>')||',' from skill s where alias is not null;

select '"'||lower(eng_name)||' link": '||jsonb_build_object('find', lower(eng_name), 'replace', '<a href="/bestiary/appendix/creatureTypes?scroll='||alias||'" class="link">'||lower(name)||'</a>')||',' from creature_type ct;

with f1 as (
select 'Flyby Attack' as a union all
select 'Lightning Reflexes' as a
)
select '"'||f1.a||' link": '||jsonb_build_object('find', f1.a, 'replace', '<a href="/feat/'||alias||'" class="link">'||upper(substring(name, 1, 1))||lower(substring(name, 2))||'</a>')||',' 
    from feat f 
        join f1 on lower(replace(f1.a, ' ', '')) = lower(f.alias)

with s1 as (
select 'cause fear' as a union all
select 'charm person' as a union all
select 'deep slumber' as a union all
select 'shout' as a
)
select '"'||s1.a||' link": '||jsonb_build_object('find', s1.a, 'replace', '<a href="/spell/'||alias||'" class="link">'||lower(name)||'</a>')||',' 
    from spell s 
        join s1 on lower(replace(s1.a, ' ', '')) = lower(s.alias);