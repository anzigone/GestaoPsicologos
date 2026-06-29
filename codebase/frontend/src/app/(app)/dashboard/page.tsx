'use client';

import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from 'recharts';
import { TrendingUp, Calendar, Users, Clock } from 'lucide-react';

const chartData = [
  { month: 'Jan', faturamento: 28000, meta: 30000 },
  { month: 'Fev', faturamento: 32000, meta: 32000 },
  { month: 'Mar', faturamento: 45000, meta: 35000 },
  { month: 'Abr', faturamento: 58000, meta: 85000 },
  { month: 'Mai', faturamento: 98000, meta: 110000 },
  { month: 'Jun', faturamento: 145680, meta: 138000 },
];

const transactions = [
  { date: '18/06/24', patient: 'Carla Mendes', professional: 'Psicólogo', service: 'Terapia Individual', value: 145, status: 'pago' },
  { date: '18/06/24', patient: 'João Silva', professional: 'Psicólogo', service: 'Terapia Individual', value: 38, status: 'pendente' },
  { date: '18/06/24', patient: 'Carla Mendes', professional: 'Psicólogo', service: 'Terapia Individual', value: 29, status: 'pago' },
  { date: '18/06/24', patient: 'Carla Mendes', professional: 'Psicólogo', service: 'Terapia Individual', value: 30, status: 'pago' },
  { date: '18/06/24', patient: 'Carla Mendes', professional: 'Psicólogo', service: 'Terapia Individual', value: 30, status: 'pago' },
  { date: '18/06/24', patient: 'João Silva', professional: 'Psicólogo', service: 'Terapia Individual', value: 38, status: 'pendente' },
  { date: '18/06/24', patient: 'Carla Mendes', professional: 'Psicólogo', service: 'Terapia Individual', value: 30, status: 'pago' },
];

const kpis = [
  { label: 'Faturamento Total (R$)', value: 'R$ 145.680,50', change: '+12.4% vs. mês ant.', icon: TrendingUp, color: 'text-teal-400' },
  { label: 'Atendimentos', value: '512', change: 'Sessões Realizadas (+8.1%)', icon: Calendar, color: 'text-blue-400' },
  { label: 'Pacientes Ativos', value: '218', change: 'Pacientes em tratamento', icon: Users, color: 'text-purple-400' },
  { label: 'A Receber', value: 'R$ 38.915,00', change: 'Total Pendente', icon: Clock, color: 'text-orange-400' },
];

export default function DashboardPage() {
  return (
    <div className="min-h-screen bg-slate-900 text-white p-6">
      <div className="max-w-6xl mx-auto">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-2xl font-bold">Dashboard Financeiro</h1>
            <p className="text-slate-400 text-sm mt-1">Visão Geral — Junho 2026</p>
          </div>
          <div className="flex items-center gap-3">
            <button className="flex items-center gap-2 bg-slate-800 border border-slate-700 text-slate-300 text-sm px-4 py-2 rounded-lg">
              <Calendar size={14} />
              01 Jan 2026 - 30 Jun 2026
            </button>
            <div className="flex items-center gap-2 bg-slate-800 border border-slate-700 px-4 py-2 rounded-lg">
              <div className="w-7 h-7 rounded-full bg-teal-600 flex items-center justify-center text-xs font-bold">A</div>
              <div>
                <p className="text-sm font-medium">Dra. Ana Silva</p>
                <p className="text-xs text-slate-400">Psicóloga</p>
              </div>
            </div>
          </div>
        </div>

        {/* KPI Cards */}
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          {kpis.map(({ label, value, change, icon: Icon, color }) => (
            <div key={label} className="bg-slate-800 rounded-xl p-5 border border-slate-700">
              <div className="flex items-center justify-between mb-3">
                <p className="text-xs text-slate-400 font-medium uppercase tracking-wider">{label}</p>
                <Icon size={16} className={color} />
              </div>
              <p className="text-2xl font-bold mb-1">{value}</p>
              <p className="text-xs text-slate-400">{change}</p>
            </div>
          ))}
        </div>

        {/* Chart + Transactions */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Area Chart */}
          <div className="bg-slate-800 rounded-xl p-6 border border-slate-700">
            <h2 className="text-sm font-semibold mb-1">Crescimento Financeiro (R$)</h2>
            <p className="text-xs text-slate-400 mb-5">Faturamento vs. Meta (Jan–Jun)</p>
            <ResponsiveContainer width="100%" height={220}>
              <AreaChart data={chartData} margin={{ top: 5, right: 10, left: 0, bottom: 0 }}>
                <defs>
                  <linearGradient id="colorFat" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#0D9488" stopOpacity={0.4} />
                    <stop offset="95%" stopColor="#0D9488" stopOpacity={0} />
                  </linearGradient>
                </defs>
                <CartesianGrid strokeDasharray="3 3" stroke="#334155" />
                <XAxis dataKey="month" tick={{ fill: '#94a3b8', fontSize: 11 }} axisLine={false} tickLine={false} />
                <YAxis
                  tick={{ fill: '#94a3b8', fontSize: 11 }}
                  axisLine={false}
                  tickLine={false}
                  tickFormatter={(v) => `${(v / 1000).toFixed(0)}k`}
                />
                <Tooltip
                  contentStyle={{ backgroundColor: '#1e293b', border: '1px solid #334155', borderRadius: 8 }}
                  labelStyle={{ color: '#e2e8f0' }}
                  formatter={(value: number) =>
                    [`R$ ${value.toLocaleString('pt-BR')}`, undefined]
                  }
                />
                <Area type="monotone" dataKey="faturamento" stroke="#0D9488" strokeWidth={2} fill="url(#colorFat)" name="Faturamento" />
                <Area type="monotone" dataKey="meta" stroke="#64748b" strokeWidth={1.5} strokeDasharray="5 5" fill="none" name="Meta" />
              </AreaChart>
            </ResponsiveContainer>
          </div>

          {/* Transactions Table */}
          <div className="bg-slate-800 rounded-xl p-6 border border-slate-700">
            <h2 className="text-sm font-semibold mb-4">Transações Recentes de Pacientes</h2>
            <div className="overflow-x-auto">
              <table className="w-full text-xs">
                <thead>
                  <tr className="text-slate-400 border-b border-slate-700">
                    {['Data', 'Paciente', 'Serviço', 'Valor', 'Status'].map((h) => (
                      <th key={h} className="text-left pb-2 pr-3 font-medium">{h}</th>
                    ))}
                  </tr>
                </thead>
                <tbody className="divide-y divide-slate-700/50">
                  {transactions.map((tx, i) => (
                    <tr key={i} className="hover:bg-slate-700/30 transition-colors">
                      <td className="py-2.5 pr-3 text-slate-400">{tx.date}</td>
                      <td className="py-2.5 pr-3 text-slate-200">{tx.patient}</td>
                      <td className="py-2.5 pr-3 text-slate-400">{tx.service}</td>
                      <td className="py-2.5 pr-3 text-slate-200">R$ {tx.value},00</td>
                      <td className="py-2.5">
                        <span
                          className={
                            tx.status === 'pago'
                              ? 'px-2 py-0.5 bg-teal-900/60 text-teal-400 rounded-full font-medium'
                              : 'px-2 py-0.5 bg-orange-900/60 text-orange-400 rounded-full font-medium'
                          }
                        >
                          {tx.status === 'pago' ? 'Pago' : 'Pendente'}
                        </span>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
