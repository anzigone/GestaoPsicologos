'use client';

import { useState, useEffect } from 'react';
import { Search, Plus, Download, Bold, Italic, List, ChevronDown } from 'lucide-react';
import { api } from '@/lib/api';
import { cn } from '@/lib/utils';
import type { Patient, FirstAnalysis, Session } from '@/types';

type Tab = 'cadastro' | 'analise' | 'evolucoes';

const ANALYSIS_SECTIONS = [
  { key: 'main_complaint', label: 'Queixa Principal' },
  { key: 'symptom_diagnosis', label: 'Sintomas / Diagnóstico' },
  { key: 'developmental_influence', label: 'Influências do Desenvolvimento' },
  { key: 'situational_issues', label: 'Questões Situacionais' },
  { key: 'biological_factors', label: 'Fatores Biológicos' },
  { key: 'strengths_resources', label: 'Forças e Recursos' },
  { key: 'addictions', label: 'Vícios / Dependências' },
  { key: 'stimuli', label: 'Estímulos / Gatilhos' },
  { key: 'thoughts', label: 'Pensamentos' },
  { key: 'behaviors', label: 'Comportamentos' },
  { key: 'affects', label: 'Afetos' },
  { key: 'physiological', label: 'Fisiológicos' },
  { key: 'treatment_goals', label: 'Objetivos do Tratamento' },
  { key: 'treatment_plan', label: 'Plano de Tratamento' },
];

export default function PatientsPage() {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [search, setSearch] = useState('');
  const [selected, setSelected] = useState<Patient | null>(null);
  const [tab, setTab] = useState<Tab>('cadastro');
  const [analysis, setAnalysis] = useState<FirstAnalysis | null>(null);
  const [sessions, setSessions] = useState<Session[]>([]);
  const [openAccordion, setOpenAccordion] = useState<string | null>('main_complaint');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    api.get<Patient[]>('/api/patients').then((data) => {
      setPatients(data);
      if (data.length > 0) setSelected(data[0]);
      setLoading(false);
    });
  }, []);

  useEffect(() => {
    if (!selected) return;
    api.get<FirstAnalysis>(`/api/patients/${selected.id}/analysis`).then(setAnalysis);
    api.get<Session[]>(`/api/patients/${selected.id}/sessions`).then(setSessions);
  }, [selected]);

  const filtered = patients.filter((p) =>
    p.name.toLowerCase().includes(search.toLowerCase())
  );

  function formatDate(iso: string) {
    return new Date(iso).toLocaleDateString('pt-BR', { day: '2-digit', month: '2-digit', year: 'numeric' });
  }

  function formatCurrency(val: number) {
    return val.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
  }

  return (
    <div className="flex h-screen overflow-hidden">
      {/* Patient list sidebar */}
      <div className="w-64 bg-white border-r border-gray-200 flex flex-col shrink-0">
        <div className="p-4 border-b border-gray-100">
          <div className="relative mb-3">
            <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
            <input
              placeholder="Buscar paciente..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="w-full pl-8 pr-3 py-2 text-xs border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-teal-500"
            />
          </div>
          <button className="w-full flex items-center justify-center gap-1.5 bg-teal-600 hover:bg-teal-700 text-white text-xs font-medium py-2 rounded-lg transition-colors">
            <Plus size={13} />
            Novo Paciente
          </button>
        </div>

        <div className="flex-1 overflow-y-auto">
          {loading ? (
            <p className="text-xs text-slate-400 text-center py-8">Carregando...</p>
          ) : (
            filtered.map((p) => (
              <button
                key={p.id}
                onClick={() => { setSelected(p); setTab('cadastro'); }}
                className={cn(
                  'w-full text-left px-4 py-3 border-b border-gray-50 transition-colors',
                  selected?.id === p.id
                    ? 'bg-teal-50 border-l-4 border-l-teal-600'
                    : 'hover:bg-gray-50'
                )}
              >
                <p className="text-sm font-medium text-slate-800">{p.name}</p>
                <p className="text-xs text-slate-400 mt-0.5">{p.phone}</p>
              </button>
            ))
          )}
        </div>
      </div>

      {/* Patient detail */}
      {selected ? (
        <div className="flex-1 flex flex-col overflow-hidden">
          {/* Header */}
          <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 rounded-full bg-teal-100 flex items-center justify-center text-teal-700 font-bold text-lg">
                {selected.name[0]}
              </div>
              <div>
                <h2 className="text-lg font-bold text-slate-800">{selected.name}</h2>
                <div className="flex items-center gap-3 mt-0.5">
                  <span className="text-sm text-slate-500">{selected.phone}</span>
                  <span className="px-2 py-0.5 bg-teal-50 text-teal-700 text-xs font-medium rounded-full">
                    Paciente Ativo
                  </span>
                </div>
              </div>
            </div>
            <a
              href={`http://localhost:8080/api/patients/${selected.id}/pdf`}
              target="_blank"
              rel="noreferrer"
              className="flex items-center gap-2 border border-gray-200 hover:bg-gray-50 text-slate-700 text-sm font-medium px-4 py-2 rounded-lg transition-colors"
            >
              <Download size={15} />
              Exportar Ficha (PDF)
            </a>
          </div>

          {/* Tabs */}
          <div className="bg-white border-b border-gray-200 px-6">
            <nav className="flex gap-6">
              {([
                { key: 'cadastro', label: 'Cadastro' },
                { key: 'analise', label: 'Primeira Análise' },
                { key: 'evolucoes', label: 'Evolução das Sessões' },
              ] as const).map(({ key, label }) => (
                <button
                  key={key}
                  onClick={() => setTab(key)}
                  className={cn(
                    'py-3 text-sm font-medium border-b-2 transition-colors -mb-px',
                    tab === key
                      ? 'border-teal-600 text-teal-700'
                      : 'border-transparent text-slate-500 hover:text-slate-700'
                  )}
                >
                  {label}
                </button>
              ))}
            </nav>
          </div>

          {/* Tab content */}
          <div className="flex-1 overflow-y-auto p-6">
            {/* Cadastro tab */}
            {tab === 'cadastro' && (
              <div className="max-w-2xl bg-white rounded-xl border border-gray-100 p-6 shadow-sm">
                <div className="grid grid-cols-2 gap-4">
                  {[
                    { label: 'Nome Completo', value: selected.name },
                    { label: 'Telefone', value: selected.phone },
                    { label: 'Data de Nascimento', value: selected.birthdate || '—' },
                    { label: 'Idade', value: selected.age ? `${selected.age} anos` : '—' },
                    { label: 'Profissão', value: selected.profession || '—' },
                    { label: 'Empresa', value: selected.company || '—' },
                    { label: 'Cidade', value: selected.city || '—' },
                    { label: 'Estado', value: selected.state || '—' },
                    { label: 'Estado Civil', value: selected.marital_status || '—' },
                  ].map(({ label, value }) => (
                    <div key={label}>
                      <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">{label}</label>
                      <input
                        defaultValue={value}
                        className="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
                      />
                    </div>
                  ))}
                  <div className="col-span-2">
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">
                      Valor da Consulta
                    </label>
                    <div className="flex items-center gap-3">
                      <input
                        defaultValue={formatCurrency(selected.consultation_fee)}
                        className="w-48 px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
                      />
                      <span className="text-xs text-slate-400">
                        Padrão: R$ 200,00 | Variação: <span className="text-teal-600 font-medium">+0%</span>
                      </span>
                    </div>
                  </div>
                </div>
                <button className="mt-5 bg-teal-600 hover:bg-teal-700 text-white text-sm font-medium px-5 py-2.5 rounded-lg transition-colors">
                  Salvar Cadastro
                </button>
              </div>
            )}

            {/* Primeira Análise tab */}
            {tab === 'analise' && analysis && (
              <div className="max-w-2xl space-y-2">
                {ANALYSIS_SECTIONS.map(({ key, label }) => (
                  <div key={key} className="bg-white rounded-xl border border-gray-100 overflow-hidden shadow-sm">
                    <button
                      onClick={() => setOpenAccordion(openAccordion === key ? null : key)}
                      className="w-full flex items-center justify-between px-5 py-3.5 text-sm font-medium text-slate-700 hover:bg-gray-50 transition-colors"
                    >
                      {label}
                      <ChevronDown
                        size={16}
                        className={cn('text-slate-400 transition-transform', openAccordion === key && 'rotate-180')}
                      />
                    </button>
                    {openAccordion === key && (
                      <div className="px-5 pb-4 border-t border-gray-100">
                        <textarea
                          defaultValue={analysis[key as keyof FirstAnalysis] as string}
                          rows={3}
                          className="w-full mt-3 px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 resize-none"
                        />
                        <button className="mt-2 text-xs bg-teal-600 hover:bg-teal-700 text-white px-3 py-1.5 rounded-md transition-colors">
                          Salvar
                        </button>
                      </div>
                    )}
                  </div>
                ))}
              </div>
            )}

            {/* Evoluções tab */}
            {tab === 'evolucoes' && (
              <div className="flex gap-6">
                <div className="flex-1 space-y-4">
                  <div className="flex items-center justify-between mb-2">
                    <h3 className="text-sm font-semibold text-slate-700">Histórico de Sessões</h3>
                    <button className="flex items-center gap-1.5 bg-teal-600 hover:bg-teal-700 text-white text-xs font-medium px-3 py-2 rounded-lg transition-colors">
                      <Plus size={13} />
                      Nova Sessão
                    </button>
                  </div>

                  {sessions.map((session) => (
                    <div key={session.id} className="bg-white rounded-xl border border-gray-100 p-4 shadow-sm">
                      <div className="flex items-center justify-between mb-2">
                        <span className="text-sm font-semibold text-slate-700">
                          Sessão — {formatDate(session.session_date)}
                        </span>
                        <span
                          className={cn(
                            'px-2 py-0.5 text-xs font-medium rounded-full',
                            session.status === 'pago'
                              ? 'bg-teal-50 text-teal-700'
                              : 'bg-orange-50 text-orange-600'
                          )}
                        >
                          {session.status === 'pago' ? 'Pago' : 'Pendente'}
                        </span>
                      </div>
                      <div className="flex gap-1 mb-2">
                        <button className="p-1.5 hover:bg-gray-100 rounded text-slate-500"><Bold size={13} /></button>
                        <button className="p-1.5 hover:bg-gray-100 rounded text-slate-500"><Italic size={13} /></button>
                        <button className="p-1.5 hover:bg-gray-100 rounded text-slate-500"><List size={13} /></button>
                      </div>
                      <textarea
                        defaultValue={session.notes}
                        rows={3}
                        placeholder="Notas da sessão..."
                        className="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 resize-none"
                      />
                      <div className="flex items-center justify-between mt-2">
                        {session.meet_link && (
                          <a
                            href={session.meet_link}
                            target="_blank"
                            rel="noreferrer"
                            className="text-xs text-teal-600 hover:underline"
                          >
                            Google Meet
                          </a>
                        )}
                        <button className="ml-auto text-xs bg-teal-600 hover:bg-teal-700 text-white px-3 py-1.5 rounded-md transition-colors">
                          Salvar Alterações
                        </button>
                      </div>
                    </div>
                  ))}
                </div>

                {/* Right panel */}
                <div className="w-56 shrink-0 space-y-4">
                  <div className="bg-white rounded-xl border border-gray-100 p-4 shadow-sm">
                    <h4 className="text-xs font-bold text-slate-500 uppercase tracking-wider mb-3">
                      Próxima Sessão
                    </h4>
                    {sessions.find((s) => s.status === 'pendente') ? (
                      <p className="text-sm text-slate-700 font-medium">
                        {formatDate(sessions.find((s) => s.status === 'pendente')!.session_date)}
                      </p>
                    ) : (
                      <p className="text-xs text-slate-400">Nenhuma agendada</p>
                    )}
                  </div>
                  <div className="bg-white rounded-xl border border-gray-100 p-4 shadow-sm">
                    <h4 className="text-xs font-bold text-slate-500 uppercase tracking-wider mb-3">
                      Resumo de Progresso
                    </h4>
                    <div className="space-y-1.5">
                      <div className="flex justify-between text-xs text-slate-600">
                        <span>Sessões realizadas</span>
                        <span className="font-medium">{sessions.filter((s) => s.status === 'pago').length}</span>
                      </div>
                      <div className="flex justify-between text-xs text-slate-600">
                        <span>Pendentes</span>
                        <span className="font-medium text-orange-500">{sessions.filter((s) => s.status === 'pendente').length}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </div>
        </div>
      ) : (
        <div className="flex-1 flex items-center justify-center text-slate-400">
          Selecione um paciente
        </div>
      )}
    </div>
  );
}
