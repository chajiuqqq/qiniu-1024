import { lusitana } from '@/app/ui/fonts';

export default function AcmeLogo() {
  return (
    <div
      className={`${lusitana.className} flex flex-row items-center leading-none text-white`}
    >
      <img src="/tv.png" alt="" className='h-12 w-12' />
      <p className="text-xl ml-4">New视频</p>
    </div>
  );
}
