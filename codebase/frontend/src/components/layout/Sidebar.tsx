'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import {
  Home,
  Calendar,
  Users,
  CircleDollarSign,
  Settings,
  LogOut,
} from 'lucide-react';
import { cn } from '@/lib/utils';

const navItems = [
  { href: '/', label: 'Início', icon: Home },
  { href: '/agenda', label: 'Agenda', icon: Calendar },
  { href: '/patients', label: 'Pacientes', icon: Users },
  { href: '/dashboard', label: 'Finanças', icon: CircleDollarSign },
  { href: '/admin/settings', label: 'Configurações', icon: Settings },
];

export default function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="flex flex-col w-60 min-h-screen bg-white border-r border-gray-200 shrink-0">
      <div className="flex items-center gap-2 px-5 py-5 border-b border-gray-100">
        <div className="w-8 h-8 rounded-full bg-teal-600 flex items-center justify-center">
          <span className="text-white text-sm font-bold">Ψ</span>
        </div>
        <span className="font-semibold text-gray-800 text-sm">PsicoGestor</span>
      </div>

      <nav className="flex-1 py-4 px-3">
        <ul className="space-y-1">
          {navItems.map(({ href, label, icon: Icon }) => {
            const active =
              href === '/' ? pathname === href : pathname.startsWith(href);
            return (
              <li key={href}>
                <Link
                  href={href}
                  className={cn(
                    'flex items-center gap-3 px-3 py-2.5 rounded-md text-sm font-medium transition-colors',
                    active
                      ? 'bg-teal-50 text-teal-700 border-l-4 border-teal-600 pl-2'
                      : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900'
                  )}
                >
                  <Icon size={18} />
                  {label}
                </Link>
              </li>
            );
          })}
        </ul>
      </nav>

      <div className="px-3 pb-4 border-t border-gray-100 pt-3">
        <button className="flex items-center gap-3 px-3 py-2.5 w-full rounded-md text-sm text-gray-500 hover:bg-gray-50 hover:text-gray-700 transition-colors">
          <LogOut size={18} />
          Sair
        </button>
      </div>
    </aside>
  );
}
