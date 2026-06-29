import Link from 'next/link';
import { Settings, Users, BarChart3, LogOut } from 'lucide-react';

const cards = [
  {
    icon: Settings,
    title: 'Minha Conta & Integrações',
    description: 'Gerenciar perfil, configurações e conexões',
    href: '/admin/settings',
  },
  {
    icon: Users,
    title: 'Pacientes & Prontuários',
    description: 'Visualizar cadastro, histórico clínico e agendamento',
    href: '/patients',
  },
  {
    icon: BarChart3,
    title: 'Finanças & Relatórios',
    description: 'Acompanhar faturamento, despesas e emitir relatórios',
    href: '/dashboard',
  },
];

export default function HomePage() {
  return (
    <div className="min-h-screen flex flex-col bg-teal-50">
      <header className="bg-slate-800 text-white px-6 py-3 flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="w-8 h-8 rounded-full bg-gray-500 flex items-center justify-center text-xs">
            DR
          </div>
          <span className="text-sm font-medium">Dra. Ana Beatriz Santos (CRP 06/123456)</span>
        </div>
        <div className="flex items-center gap-4">
          <button className="flex items-center gap-1.5 text-sm bg-teal-600 hover:bg-teal-700 px-3 py-1.5 rounded transition-colors">
            <LogOut size={14} />
            Sair
          </button>
          <div className="flex items-center gap-1.5">
            <div className="w-6 h-6 rounded-full bg-teal-500 flex items-center justify-center">
              <span className="text-white text-xs font-bold">Ψ</span>
            </div>
            <span className="text-sm font-semibold">PsicoGestor</span>
          </div>
        </div>
      </header>

      <div className="flex-1 flex flex-col items-center justify-center px-6 py-12">
        <h1 className="text-4xl font-bold text-slate-800 text-center mb-2">
          Olá, Dra. Ana!
        </h1>
        <h2 className="text-3xl font-bold text-slate-700 text-center mb-3">
          O que deseja gerenciar hoje?
        </h2>
        <p className="text-slate-500 text-center mb-10">
          Seja bem-vindo(a) ao seu painel de controle. Acesse as ferramentas abaixo:
        </p>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 w-full max-w-4xl">
          {cards.map(({ icon: Icon, title, description, href }) => (
            <div
              key={href}
              className="bg-white border border-teal-100 rounded-xl p-8 flex flex-col items-center text-center shadow-sm hover:shadow-md transition-shadow"
            >
              <div className="mb-4 text-teal-600">
                <Icon size={40} strokeWidth={1.5} />
              </div>
              <h3 className="font-bold text-slate-800 text-lg mb-2">{title}</h3>
              <p className="text-slate-500 text-sm mb-6">{description}</p>
              <Link
                href={href}
                className="px-6 py-2 border border-slate-300 rounded-md text-sm text-slate-600 hover:bg-slate-50 transition-colors"
              >
                Acessar
              </Link>
            </div>
          ))}
        </div>
      </div>

      <footer className="py-4 text-center text-xs text-slate-400">
        PsicoGestor © 2026 | CRP 06/123456 | Termos de Uso | Ajuda
      </footer>
    </div>
  );
}
