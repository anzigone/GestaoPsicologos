'use client';

import { useState, useEffect, useCallback, useRef } from 'react';
import { Search, Plus, Download, Bold, Italic, List, ChevronDown, X, CheckCircle, XCircle, Trash2, UserX, UserCheck, EyeOff } from 'lucide-react';
import { api } from '@/lib/api';
import { cn } from '@/lib/utils';
import type { Patient, FirstAnalysis, Session, User } from '@/types';

type Tab = 'cadastro' | 'analise' | 'evolucoes';
type ToastState = { message: string; type: 'success' | 'error' } | null;

const ANALYSIS_SECTIONS: { key: keyof Omit<FirstAnalysis, 'patient_id' | 'updated_at'>; label: string }[] = [
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

const emptyAnalysis = (): Partial<FirstAnalysis> => ({
  main_complaint: '', symptom_diagnosis: '', developmental_influence: '',
  situational_issues: '', biological_factors: '', strengths_resources: '',
  addictions: '', stimuli: '', thoughts: '', behaviors: '',
  affects: '', physiological: '', treatment_goals: '', treatment_plan: '',
});

export default function PatientsPage() {
  const [baseFee, setBaseFee] = useState(0);

  const [patients, setPatients] = useState<Patient[]>([]);
  const [search, setSearch] = useState('');
  const [loadingList, setLoadingList] = useState(true);

  const [selected, setSelected] = useState<Patient | null>(null);
  const [tab, setTab] = useState<Tab>('cadastro');

  // Cadastro form (controlled)
  const [form, setForm] = useState<Partial<Patient>>({});
  const [savingForm, setSavingForm] = useState(false);

  // New patient modal
  const [showModal, setShowModal] = useState(false);
  const [newForm, setNewForm] = useState({ name: '', phone: '', consultation_fee: '' });
  const [savingNew, setSavingNew] = useState(false);

  // Analysis
  const [analysis, setAnalysis] = useState<Partial<FirstAnalysis>>(emptyAnalysis());
  const [openAccordion, setOpenAccordion] = useState<string | null>('main_complaint');
  const [savingAnalysis, setSavingAnalysis] = useState(false);

  const [sessions, setSessions] = useState<Session[]>([]);

  // Delete confirmation
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [deletingPatient, setDeletingPatient] = useState(false);
  const [togglingActive, setTogglingActive] = useState(false);

  // Show inactive patients in sidebar
  const [showInactive, setShowInactive] = useState(false);

  const [toast, setToast] = useState<ToastState>(null);
  const showToast = useCallback((message: string, type: 'success' | 'error' = 'success') => {
    setToast({ message, type });
    setTimeout(() => setToast(null), 3000);
  }, []);

  // Load psychologist base_fee for variation calculation
  useEffect(() => {
    api.get<User>('/api/psychologist').then((u) => setBaseFee(u.base_fee ?? 0));
  }, []);

  // Debounced patient list search
  const debounceRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  useEffect(() => {
    if (debounceRef.current) clearTimeout(debounceRef.current);
    debounceRef.current = setTimeout(() => {
      setLoadingList(true);
      const inactiveParam = showInactive ? '&include_inactive=1' : '';
      api.get<Patient[]>(`/api/patients?q=${encodeURIComponent(search)}${inactiveParam}`)
        .then((data) => {
          setPatients(data);
          if (!selected && data.length > 0) setSelected(data[0]);
        })
        .catch(() => {})
        .finally(() => setLoadingList(false));
    }, 300);
    return () => { if (debounceRef.current) clearTimeout(debounceRef.current); };
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [search, showInactive]);

  // Load patient data on selection
  useEffect(() => {
    if (!selected) return;
    setForm({ ...selected });
    setAnalysis(emptyAnalysis());
    api.get<FirstAnalysis>(`/api/patients/${selected.id}/analysis`)
      .then((a) => setAnalysis({ ...emptyAnalysis(), ...a }))
      .catch(() => {});
    api.get<Session[]>(`/api/patients/${selected.id}/sessions`)
      .then(setSessions)
      .catch(() => setSessions([]));
  }, [selected]);

  function feeVariation(fee: number): string {
    if (!baseFee || baseFee === 0) return '';
    const diff = ((fee - baseFee) / baseFee) * 100;
    if (Math.abs(diff) < 0.01) return '= valor padrão';
    return diff > 0
      ? `+${diff.toFixed(0)}% acima do valor padrão`
      : `${diff.toFixed(0)}% abaixo do valor padrão`;
  }

  function feeVariationColor(fee: number): string {
    if (!baseFee) return 'text-slate-400';
    const diff = ((fee - baseFee) / baseFee) * 100;
    if (Math.abs(diff) < 0.01) return 'text-slate-400';
    return diff > 0 ? 'text-teal-600' : 'text-orange-500';
  }

  async function handleSaveForm(e: React.FormEvent) {
    e.preventDefault();
    if (!selected) return;
    setSavingForm(true);
    try {
      const updated = await api.put<Patient>(`/api/patients/${selected.id}`, {
        name: form.name,
        phone: form.phone,
        birthdate: form.birthdate,
        age: Number(form.age) || 0,
        profession: form.profession,
        company: form.company,
        city: form.city,
        state: form.state,
        marital_status: form.marital_status,
        consultation_fee: Number(form.consultation_fee) || 0,
      });
      setSelected(updated);
      setPatients((ps) => ps.map((p) => (p.id === updated.id ? updated : p)));
      showToast('Cadastro salvo com sucesso!');
    } catch {
      showToast('Erro ao salvar cadastro.', 'error');
    } finally {
      setSavingForm(false);
    }
  }

  async function handleCreatePatient(e: React.FormEvent) {
    e.preventDefault();
    setSavingNew(true);
    try {
      const created = await api.post<Patient>('/api/patients', {
        name: newForm.name,
        phone: newForm.phone,
        consultation_fee: parseFloat(newForm.consultation_fee) || 0,
      });
      setPatients((ps) => [...ps, created].sort((a, b) => a.name.localeCompare(b.name)));
      setSelected(created);
      setTab('cadastro');
      setShowModal(false);
      setNewForm({ name: '', phone: '', consultation_fee: '' });
      showToast('Paciente cadastrado com sucesso!');
    } catch {
      showToast('Erro ao cadastrar paciente.', 'error');
    } finally {
      setSavingNew(false);
    }
  }

  async function handleDeletePatient() {
    if (!selected) return;
    setDeletingPatient(true);
    try {
      await api.delete(`/api/patients/${selected.id}`);
      const remaining = patients.filter((p) => p.id !== selected.id);
      setPatients(remaining);
      setSelected(remaining.length > 0 ? remaining[0] : null);
      setShowDeleteConfirm(false);
      showToast('Paciente removido com sucesso!');
    } catch {
      showToast('Erro ao remover paciente.', 'error');
    } finally {
      setDeletingPatient(false);
    }
  }

  async function handleToggleActive() {
    if (!selected) return;
    setTogglingActive(true);
    try {
      const updated = await api.patch<Patient>(`/api/patients/${selected.id}`, { active: !selected.active });
      if (!updated.active && !showInactive) {
        const remaining = patients.filter((p) => p.id !== selected.id);
        setPatients(remaining);
        setSelected(remaining.length > 0 ? remaining[0] : null);
      } else {
        setSelected(updated);
        setPatients((ps) => ps.map((p) => (p.id === updated.id ? updated : p)));
      }
      showToast(updated.active ? 'Paciente reativado com sucesso!' : 'Paciente desativado.');
    } catch {
      showToast('Erro ao alterar status do paciente.', 'error');
    } finally {
      setTogglingActive(false);
    }
  }

  async function handleSaveAnalysis() {
    if (!selected) return;
    setSavingAnalysis(true);
    try {
      const saved = await api.put<FirstAnalysis>(`/api/patients/${selected.id}/analysis`, analysis);
      setAnalysis({ ...emptyAnalysis(), ...saved });
      showToast('Primeira Análise salva com sucesso!');
    } catch {
      showToast('Erro ao salvar análise.', 'error');
    } finally {
      setSavingAnalysis(false);
    }
  }

  function formatDate(iso: string) {
    return new Date(iso).toLocaleDateString('pt-BR', { day: '2-digit', month: '2-digit', year: 'numeric' });
  }

  const inputClass = 'w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500';

  return (
    <div className="flex h-screen overflow-hidden">
      {/* Toast */}
      {toast && (
        <div className={`fixed top-5 right-5 z-50 flex items-center gap-2 px-4 py-3 rounded-lg shadow-lg text-white text-sm font-medium ${toast.type === 'success' ? 'bg-teal-600' : 'bg-red-500'}`}>
          {toast.type === 'success' ? <CheckCircle size={16} /> : <XCircle size={16} />}
          {toast.message}
        </div>
      )}

      {/* New patient modal */}
      {showModal && (
        <div className="fixed inset-0 z-40 bg-black/40 flex items-center justify-center">
          <form
            onSubmit={handleCreatePatient}
            className="bg-white rounded-xl shadow-xl w-full max-w-md p-6"
          >
            <div className="flex items-center justify-between mb-5">
              <h2 className="text-base font-bold text-slate-800">Novo Paciente</h2>
              <button type="button" onClick={() => setShowModal(false)} className="text-slate-400 hover:text-slate-600">
                <X size={18} />
              </button>
            </div>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Nome Completo *</label>
                <input
                  value={newForm.name}
                  onChange={(e) => setNewForm((f) => ({ ...f, name: e.target.value }))}
                  required
                  placeholder="Ex: Carlos Drummond"
                  className={inputClass}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Telefone</label>
                <input
                  value={newForm.phone}
                  onChange={(e) => setNewForm((f) => ({ ...f, phone: e.target.value }))}
                  placeholder="(21) 99999-8888"
                  className={inputClass}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">
                  Valor da Consulta (R$)
                  {baseFee > 0 && (
                    <span className="ml-2 text-xs font-normal text-slate-400">
                      padrão: R$ {baseFee.toFixed(2).replace('.', ',')}
                    </span>
                  )}
                </label>
                <input
                  type="number"
                  min="0"
                  step="0.01"
                  value={newForm.consultation_fee}
                  onChange={(e) => setNewForm((f) => ({ ...f, consultation_fee: e.target.value }))}
                  placeholder="0,00"
                  className={inputClass}
                />
                {newForm.consultation_fee && baseFee > 0 && (
                  <p className={`text-xs mt-1 ${feeVariationColor(parseFloat(newForm.consultation_fee) || 0)}`}>
                    {feeVariation(parseFloat(newForm.consultation_fee) || 0)}
                  </p>
                )}
              </div>
            </div>
            <div className="flex gap-3 mt-6">
              <button
                type="button"
                onClick={() => setShowModal(false)}
                className="flex-1 border border-gray-200 text-slate-700 text-sm font-medium py-2.5 rounded-lg hover:bg-gray-50 transition-colors"
              >
                Cancelar
              </button>
              <button
                type="submit"
                disabled={savingNew}
                className="flex-1 bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white text-sm font-medium py-2.5 rounded-lg transition-colors"
              >
                {savingNew ? 'Salvando...' : 'Salvar Novo Paciente'}
              </button>
            </div>
          </form>
        </div>
      )}

      {/* Delete confirmation modal */}
      {showDeleteConfirm && selected && (
        <div className="fixed inset-0 z-40 bg-black/40 flex items-center justify-center">
          <div className="bg-white rounded-xl shadow-xl w-full max-w-sm p-6">
            <div className="flex items-center gap-3 mb-3">
              <div className="w-10 h-10 rounded-full bg-red-100 flex items-center justify-center">
                <Trash2 size={18} className="text-red-600" />
              </div>
              <h2 className="text-base font-bold text-slate-800">Remover Paciente</h2>
            </div>
            <p className="text-sm text-slate-600 mb-5">
              Tem certeza que deseja remover <strong>{selected.name}</strong>? Todos os dados do paciente serão perdidos permanentemente.
            </p>
            <div className="flex gap-3">
              <button
                onClick={() => setShowDeleteConfirm(false)}
                className="flex-1 border border-gray-200 text-slate-700 text-sm font-medium py-2.5 rounded-lg hover:bg-gray-50 transition-colors"
              >
                Cancelar
              </button>
              <button
                onClick={handleDeletePatient}
                disabled={deletingPatient}
                className="flex-1 bg-red-600 hover:bg-red-700 disabled:opacity-60 text-white text-sm font-medium py-2.5 rounded-lg transition-colors"
              >
                {deletingPatient ? 'Removendo...' : 'Sim, Remover'}
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Sidebar */}
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
          <button
            onClick={() => setShowModal(true)}
            className="w-full flex items-center justify-center gap-1.5 bg-teal-600 hover:bg-teal-700 text-white text-xs font-medium py-2 rounded-lg transition-colors"
          >
            <Plus size={13} />
            Novo Paciente
          </button>
          <button
            onClick={() => setShowInactive((v) => !v)}
            className={cn(
              'w-full flex items-center justify-center gap-1.5 mt-2 text-xs font-medium py-1.5 rounded-lg transition-colors border',
              showInactive
                ? 'bg-slate-100 border-slate-200 text-slate-700'
                : 'border-transparent text-slate-400 hover:text-slate-600 hover:bg-gray-50',
            )}
          >
            <EyeOff size={12} />
            {showInactive ? 'Ocultar inativos' : 'Mostrar inativos'}
          </button>
        </div>

        <div className="flex-1 overflow-y-auto">
          {loadingList ? (
            <p className="text-xs text-slate-400 text-center py-8">Carregando...</p>
          ) : patients.length === 0 ? (
            <p className="text-xs text-slate-400 text-center py-8">Nenhum paciente encontrado</p>
          ) : (
            patients.map((p) => (
              <button
                key={p.id}
                onClick={() => { setSelected(p); setTab('cadastro'); }}
                className={cn(
                  'w-full text-left px-4 py-3 border-b border-gray-50 transition-colors',
                  selected?.id === p.id
                    ? 'bg-teal-50 border-l-4 border-l-teal-600'
                    : 'hover:bg-gray-50',
                  !p.active && 'opacity-60',
                )}
              >
                <p className="text-sm font-medium text-slate-800">{p.name}</p>
                <div className="flex items-center gap-2 mt-0.5">
                  <p className="text-xs text-slate-400">{p.phone}</p>
                  {!p.active && (
                    <span className="text-xs text-red-400 font-medium">Inativo</span>
                  )}
                </div>
              </button>
            ))
          )}
        </div>
      </div>

      {/* Detail panel */}
      {selected ? (
        <div className="flex-1 flex flex-col overflow-hidden">
          {/* Header */}
          <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between">
            <div className="flex items-center gap-4">
              <div className={cn(
                'w-12 h-12 rounded-full flex items-center justify-center font-bold text-lg',
                selected.active ? 'bg-teal-100 text-teal-700' : 'bg-slate-100 text-slate-500',
              )}>
                {selected.name[0]}
              </div>
              <div>
                <h2 className="text-lg font-bold text-slate-800">{selected.name}</h2>
                <div className="flex items-center gap-3 mt-0.5">
                  <span className="text-sm text-slate-500">{selected.phone}</span>
                  <span className={cn(
                    'px-2 py-0.5 text-xs font-medium rounded-full',
                    selected.active ? 'bg-teal-50 text-teal-700' : 'bg-red-50 text-red-600',
                  )}>
                    {selected.active ? 'Paciente Ativo' : 'Paciente Inativo'}
                  </span>
                </div>
              </div>
            </div>
            <div className="flex items-center gap-2">
              <button
                onClick={handleToggleActive}
                disabled={togglingActive}
                className={cn(
                  'flex items-center gap-2 border text-sm font-medium px-4 py-2 rounded-lg transition-colors disabled:opacity-60',
                  selected.active
                    ? 'border-orange-200 text-orange-600 hover:bg-orange-50'
                    : 'border-teal-200 text-teal-600 hover:bg-teal-50',
                )}
              >
                {selected.active
                  ? <><UserX size={15} /> Desativar</>
                  : <><UserCheck size={15} /> Reativar</>}
              </button>
              <button
                onClick={() => setShowDeleteConfirm(true)}
                className="flex items-center gap-2 border border-red-200 hover:bg-red-50 text-red-600 text-sm font-medium px-4 py-2 rounded-lg transition-colors"
              >
                <Trash2 size={15} />
                Remover
              </button>
              <a
                href={`http://localhost:8080/api/patients/${selected.id}/pdf`}
                target="_blank"
                rel="noreferrer"
                className="flex items-center gap-2 border border-gray-200 hover:bg-gray-50 text-slate-700 text-sm font-medium px-4 py-2 rounded-lg transition-colors"
              >
                <Download size={15} />
                Exportar PDF
              </a>
            </div>
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
                      : 'border-transparent text-slate-500 hover:text-slate-700',
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
              <form onSubmit={handleSaveForm} className="max-w-2xl bg-white rounded-xl border border-gray-100 p-6 shadow-sm">
                <div className="grid grid-cols-2 gap-4">
                  <div className="col-span-2">
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Nome Completo *</label>
                    <input
                      value={form.name ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, name: e.target.value }))}
                      required
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Telefone</label>
                    <input
                      value={form.phone ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, phone: e.target.value }))}
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Data de Nascimento</label>
                    <input
                      value={form.birthdate ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, birthdate: e.target.value }))}
                      placeholder="AAAA-MM-DD"
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Idade</label>
                    <input
                      type="number"
                      min="0"
                      value={form.age ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, age: parseInt(e.target.value) || 0 }))}
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Estado Civil</label>
                    <input
                      value={form.marital_status ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, marital_status: e.target.value }))}
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Profissão</label>
                    <input
                      value={form.profession ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, profession: e.target.value }))}
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Empresa</label>
                    <input
                      value={form.company ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, company: e.target.value }))}
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Cidade</label>
                    <input
                      value={form.city ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, city: e.target.value }))}
                      className={inputClass}
                    />
                  </div>
                  <div>
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">Estado (UF)</label>
                    <input
                      value={form.state ?? ''}
                      onChange={(e) => setForm((f) => ({ ...f, state: e.target.value }))}
                      maxLength={2}
                      placeholder="SP"
                      className={inputClass}
                    />
                  </div>
                  <div className="col-span-2">
                    <label className="block text-xs font-medium text-slate-400 uppercase tracking-wider mb-1">
                      Valor da Consulta (R$)
                    </label>
                    <div className="flex items-center gap-3">
                      <input
                        type="number"
                        min="0"
                        step="0.01"
                        value={form.consultation_fee ?? 0}
                        onChange={(e) => setForm((f) => ({ ...f, consultation_fee: parseFloat(e.target.value) || 0 }))}
                        className="w-48 px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
                      />
                      {baseFee > 0 && (
                        <span className={`text-xs font-medium ${feeVariationColor(form.consultation_fee ?? 0)}`}>
                          {feeVariation(form.consultation_fee ?? 0)}
                        </span>
                      )}
                    </div>
                  </div>
                </div>
                <button
                  type="submit"
                  disabled={savingForm}
                  className="mt-5 bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white text-sm font-medium px-5 py-2.5 rounded-lg transition-colors"
                >
                  {savingForm ? 'Salvando...' : 'Salvar Cadastro'}
                </button>
              </form>
            )}

            {/* Primeira Análise tab */}
            {tab === 'analise' && (
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
                          value={(analysis[key] as string) ?? ''}
                          onChange={(e) => setAnalysis((a) => ({ ...a, [key]: e.target.value }))}
                          rows={4}
                          className="w-full mt-3 px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 resize-none"
                        />
                      </div>
                    )}
                  </div>
                ))}
                <div className="pt-2">
                  <button
                    onClick={handleSaveAnalysis}
                    disabled={savingAnalysis}
                    className="bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white text-sm font-semibold px-6 py-2.5 rounded-lg transition-colors"
                  >
                    {savingAnalysis ? 'Salvando...' : 'Salvar Primeira Análise'}
                  </button>
                </div>
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

                  {sessions.length === 0 ? (
                    <p className="text-sm text-slate-400">Nenhuma sessão registrada.</p>
                  ) : sessions.map((session) => (
                    <div key={session.id} className="bg-white rounded-xl border border-gray-100 p-4 shadow-sm">
                      <div className="flex items-center justify-between mb-2">
                        <span className="text-sm font-semibold text-slate-700">
                          Sessão — {formatDate(session.session_date)}
                        </span>
                        <span className={cn(
                          'px-2 py-0.5 text-xs font-medium rounded-full',
                          session.status === 'pago' ? 'bg-teal-50 text-teal-700' : 'bg-orange-50 text-orange-600',
                        )}>
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
                          <a href={session.meet_link} target="_blank" rel="noreferrer" className="text-xs text-teal-600 hover:underline">
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

                <div className="w-56 shrink-0 space-y-4">
                  <div className="bg-white rounded-xl border border-gray-100 p-4 shadow-sm">
                    <h4 className="text-xs font-bold text-slate-500 uppercase tracking-wider mb-3">Próxima Sessão</h4>
                    {sessions.find((s) => s.status === 'pendente') ? (
                      <p className="text-sm text-slate-700 font-medium">
                        {formatDate(sessions.find((s) => s.status === 'pendente')!.session_date)}
                      </p>
                    ) : (
                      <p className="text-xs text-slate-400">Nenhuma agendada</p>
                    )}
                  </div>
                  <div className="bg-white rounded-xl border border-gray-100 p-4 shadow-sm">
                    <h4 className="text-xs font-bold text-slate-500 uppercase tracking-wider mb-3">Resumo de Progresso</h4>
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
          Selecione um paciente ou cadastre um novo
        </div>
      )}
    </div>
  );
}
