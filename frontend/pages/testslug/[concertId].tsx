import { useRouter } from 'next/router';
import { useEffect } from 'react';

const testslug = () => {
    const router = useRouter();
    console.log("router.query in test = " , router.query);
    useEffect(()=>{
        console.log(router);

    },[]);
  return (
    <div>[index]</div>
  )
}

export default testslug