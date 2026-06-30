'use client';

import { useState, useEffect } from 'react';
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
import { api } from '@/lib/api';
import type { DashboardStats, ChartPoint, Transaction } from '@/types';

const EMPTY_STATS: DashboardStats = { total_revenue: 0, total_sessions: 0, active_patients: 0, pending_amount: 0 };

function brl(value: number) {
  return value.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
}

export default function DashboardPage() {
  const [stats, setStats] = useState<DashboardStats>(EMPTY_STATS);
  const [chartData, setChartData] = useState<ChartPoint[]>([]);
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [statusFilter, setStatusFilter] = useState<'' | 'pago' | 'pendente'>('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Promise.all([
      api.get<DashboardStats>('/api/dashboard/stats'),
      api.get<ChartPoint[]>('/api/dashboard/charts'),
    ]).then(([s, c]) => {
      setStats(s);
      setChartData(c);
    }).finally(() => setLoading(false));
  }, []);

  useEffect(() => {
    const qs = statusFilter ? `?status=${statusFilter}` : '';
    api.get<Transaction[]>(`/api/dashboard/transactions${qs}`)
      .then(setTransactions)
      .catch(() => setTransactions([]));
  }, [statusFilter]);

  const kpis = [
    { label: 'Faturamento Total', value: loading ? '—' : brl(stats.total_revenue), change: 'Sessões pagas', icon: TrendingUp, color: 'text-teal-400' },
    { label: 'Atendimentos', value: loading ? '—' : String(stats.total_sessions), change: 'Total de sessões', icon: Calendar, color: 'text-blue-400' },
    { label: 'Pacientes Ativos', value: loading ? '—' : String(stats.active_patients), change: 'Em tratamento', icon: Users, color: 'text-purple-400' },
    { label: 'A Receber', value: loading ? '—' : brl(stats.pending_amount), change: 'Total pendente', icon: Clock, color: 'text-orange-400' },
  ];

  return (
    <div className="min-h-screen bg-slate-900 text-white p-6">
      <div className="max-w-6xl mx-auto">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-2xl font-bold">Dashboard Financeiro</h1>
            <p className="text-slate-400 text-sm mt-1">Visão Geral Acumulada</p>
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
            <h2 className="text-sm font-semibold mb-1">Faturamento Mensal (R$)</h2>
            <p className="text-xs text-slate-400 mb-5">Sessões pagas — últimos 6 meses</p>
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
                  tickFormatter={(v: number) => `${(v / 1000).toFixed(0)}k`}
                />
                <Tooltip
                  contentStyle={{ backgroundColor: '#1e293b', border: '1px solid #334155', borderRadius: 8 }}
                  labelStyle={{ color: '#e2e8f0' }}
                  formatter={(value: number) => [`R$ ${value.toLocaleString('pt-BR')}`, 'Faturamento']}
                />
                <Area type="monotone" dataKey="faturamento" stroke="#0D9488" strokeWidth={2} fill="url(#colorFat)" name="Faturamento" />
              </AreaChart>
            </ResponsiveContainer>
          </div>

          {/* Transactions Table */}
          <div className="bg-slate-800 rounded-xl p-6 border border-slate-700">
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-sm font-semibold">Transações Recentes</h2>
              <div className="flex gap-1">
                {(['', 'pago', 'pendente'] as const).map((f) => (
                  <button
                    key={f}
                    onClick={() => setStatusFilter(f)}
                    className={`px-2.5 py-1 text-xs rounded-md font-medium transition-colors ${
                      statusFilter === f
                        ? 'bg-teal-600 text-white'
                        : 'bg-slate-700 text-slate-300 hover:bg-slate-600'
                    }`}
                  >
                    {f === '' ? 'Todos' : f === 'pago' ? 'Pagos' : 'Pendentes'}
                  </button>
                ))}
              </div>
            </div>
            <div className="overflow-x-auto">
              {transactions.length === 0 ? (
                <p className="text-xs text-slate-400 py-4 text-center">Nenhuma transação encontrada.</p>
              ) : (
                <table className="w-full text-xs">
                  <thead>
                    <tr className="text-slate-400 border-b border-slate-700">
                      {['Data', 'Paciente', 'Valor', 'Status'].map((h) => (
                        <th key={h} className="text-left pb-2 pr-3 font-medium">{h}</th>
                      ))}
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-slate-700/50">
                    {transactions.map((tx) => (
                      <tr key={tx.id} className="hover:bg-slate-700/30 transition-colors">
                        <td className="py-2.5 pr-3 text-slate-400">{tx.date}</td>
                        <td className="py-2.5 pr-3 text-slate-200">{tx.patient_name}</td>
                        <td className="py-2.5 pr-3 text-slate-200">{brl(tx.value)}</td>
                        <td className="py-2.5">
                          <span className={tx.status === 'pago'
                            ? 'px-2 py-0.5 bg-teal-900/60 text-teal-400 rounded-full font-medium'
                            : 'px-2 py-0.5 bg-orange-900/60 text-orange-400 rounded-full font-medium'
                          }>
                            {tx.status === 'pago' ? 'Pago' : 'Pendente'}
                          </span>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
