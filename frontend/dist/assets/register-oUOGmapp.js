import{r as N,u as v,a as d,j as e,I as u,B as p,V as n,L as C,E as h,e as E,p as g}from"./index-BsUQLVsL.js";const b="_register_psekc_1",T="_header_psekc_11",I="_check_psekc_32",a={register:b,header:T,check:I};function R(){const[o,l]=N.useState(!1),m=v(),{userValue:r,getUserValue:x,isValid:j}=d(E),{userValue:t,getUserValue:w,isValid:f}=d(g),{userValue:i,getUserValue:V}=d(g);async function k(c){c.preventDefault();try{(await fetch(h+"/auth/signup",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({email:r,password:t})})).ok&&m("/")}catch(s){if(s instanceof Error)throw console.error(s.message),new Error("회원가입에 실패했습니다.")}}async function y(c){c.preventDefault();try{const s=await fetch(`${h}/auth/email/check?email=${r}`,{method:"GET",headers:{"Content-Type":"application/json"}});s.status===400?l(!1):l(!0),s.ok&&l(!1)}catch{throw new Error("중복 체크에 실패했습니다.")}}const _=o||t!==i;return e.jsxs(e.Fragment,{children:[e.jsx("header",{className:a.header,children:e.jsx("h1",{children:"회원가입"})}),e.jsxs("form",{className:a.register,onSubmit:k,children:[e.jsxs("div",{children:[e.jsx(u,{id:"email",type:"email",label:"이메일",className:a.login__input,value:r,onChange:x}),e.jsx(p,{className:a.check,type:"button",onClick:y,children:"중복 체크"})]}),e.jsx(n,{condition:!o,type:"success",message:"사용할 수 있는 이메일입니다."}),e.jsx(n,{condition:j&&r==="",type:"warning",message:"이메일이 유효하지 않습니다."}),e.jsx(n,{condition:o,type:"warning",message:"중복된 이메일입니다."}),e.jsx(u,{id:"password",type:"password",label:"패스워드",className:a.login__input,value:t,onChange:w}),e.jsx(u,{id:"password-check",type:"password",label:"패스워드",className:a.login__input,value:i,onChange:V}),e.jsx(n,{condition:t.length>0&&!f,type:"warning",message:"암호가 유효하지 않습니다."}),e.jsx(n,{condition:i.length>0&&t!==i,type:"warning",message:"암호가 일치하지 않습니다."}),e.jsx(p,{disabled:_,children:"회원가입"}),e.jsx("nav",{className:a.nav,children:e.jsx(C,{to:"/",span:"이미 회원이신가요?",strong:"로그인하기"})})]})]})}export{R as default};