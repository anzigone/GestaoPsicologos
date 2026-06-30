'use client';

import { useState, useEffect, useCallback } from 'react';
import { CheckCircle, XCircle, Link as LinkIcon } from 'lucide-react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { api } from '@/lib/api';
import type { User } from '@/types';

type ToastState = { message: string; type: 'success' | 'error' } | null;

function useToast() {
  const [toast, setToast] = useState<ToastState>(null);
  const show = useCallback((message: string, type: 'success' | 'error' = 'success') => {
    setToast({ message, type });
    setTimeout(() => setToast(null), 3000);
  }, []);
  return { toast, show };
}

const profileSchema = z.object({
  name: z.string().min(2, 'Nome deve ter ao menos 2 caracteres'),
  crp: z.string().regex(/^\d{2}\/\d{1,6}$/, 'CRP inválido (ex: 06/123456)').or(z.literal('')),
  specialty: z.string().optional(),
  phone: z.string().optional(),
  base_fee: z.coerce.number().min(0, 'Valor não pode ser negativo'),
});

const passwordSchema = z.object({
  current_password: z.string().min(1, 'Informe a senha atual'),
  new_password: z.string().min(6, 'A nova senha deve ter ao menos 6 caracteres'),
  confirm_password: z.string().min(1, 'Confirme a nova senha'),
}).refine((d) => d.new_password === d.confirm_password, {
  message: 'As senhas não coincidem',
  path: ['confirm_password'],
});

type ProfileData = z.infer<typeof profileSchema>;
type PasswordData = z.infer<typeof passwordSchema>;

export default function AdminSettingsPage() {
  const { toast, show } = useToast();
  const [loadingProfile, setLoadingProfile] = useState(true);
  const [googleConnected] = useState(true);
  const [outlookConnected] = useState(false);

  const profileForm = useForm<ProfileData>({ resolver: zodResolver(profileSchema) });
  const passwordForm = useForm<PasswordData>({ resolver: zodResolver(passwordSchema) });

  useEffect(() => {
    api.get<User>('/api/psychologist')
      .then((u) => {
        profileForm.reset({
          name: u.name,
          crp: u.crp ?? '',
          specialty: u.specialty ?? '',
          phone: u.phone ?? '',
          base_fee: u.base_fee ?? 0,
        });
      })
      .catch(() => show('Erro ao carregar dados do perfil.', 'error'))
      .finally(() => setLoadingProfile(false));
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  async function onSaveProfile(data: ProfileData) {
    try {
      await api.put<User>('/api/psychologist', data);
      show('Dados salvos com sucesso!');
    } catch {
      show('Erro ao salvar dados. Tente novamente.', 'error');
    }
  }

  async function onChangePassword(data: PasswordData) {
    try {
      await api.post('/api/auth/change-password', {
        current_password: data.current_password,
        new_password: data.new_password,
      });
      passwordForm.reset();
      show('Senha alterada com sucesso!');
    } catch {
      passwordForm.setError('current_password', { message: 'Senha atual incorreta' });
    }
  }

  const inputClass = 'w-full px-3 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500';
  const errorClass = 'text-red-500 text-xs mt-1';

  return (
    <div className="p-6 max-w-5xl mx-auto">
      {toast && (
        <div className={`fixed top-5 right-5 z-50 flex items-center gap-2 px-4 py-3 rounded-lg shadow-lg text-white text-sm font-medium transition-all ${
          toast.type === 'success' ? 'bg-teal-600' : 'bg-red-500'
        }`}>
          {toast.type === 'success' ? <CheckCircle size={16} /> : <XCircle size={16} />}
          {toast.message}
        </div>
      )}

      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold text-slate-800">Configurações da Conta</h1>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Left column */}
        <div className="space-y-6">
          {/* Profile form */}
          <form onSubmit={profileForm.handleSubmit(onSaveProfile)} className="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h2 className="text-xs font-bold text-slate-500 uppercase tracking-widest mb-5">
              Dados do Psicólogo
            </h2>

            {loadingProfile ? (
              <div className="space-y-3 py-2">
                {[...Array(5)].map((_, i) => (
                  <div key={i} className="h-10 bg-gray-100 rounded animate-pulse" />
                ))}
              </div>
            ) : (
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-1">Nome Completo</label>
                  <input {...profileForm.register('name')} className={inputClass} />
                  {profileForm.formState.errors.name && (
                    <p className={errorClass}>{profileForm.formState.errors.name.message}</p>
                  )}
                </div>
                <div className="grid grid-cols-2 gap-3">
                  <div>
                    <label className="block text-sm font-medium text-slate-700 mb-1">CRP</label>
                    <input {...profileForm.register('crp')} placeholder="06/000000" className={inputClass} />
                    {profileForm.formState.errors.crp && (
                      <p className={errorClass}>{profileForm.formState.errors.crp.message}</p>
                    )}
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-slate-700 mb-1">Telefone</label>
                    <input {...profileForm.register('phone')} placeholder="(11) 99999-9999" className={inputClass} />
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-1">Especialidade</label>
                  <input {...profileForm.register('specialty')} placeholder="Ex: Terapia Cognitivo Comportamental" className={inputClass} />
                </div>
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-1">
                    Valor Padrão da Consulta (R$)
                  </label>
                  <input
                    type="number"
                    min="0"
                    step="0.01"
                    {...profileForm.register('base_fee')}
                    placeholder="0,00"
                    className={inputClass}
                  />
                  {profileForm.formState.errors.base_fee && (
                    <p className={errorClass}>{profileForm.formState.errors.base_fee.message}</p>
                  )}
                </div>
              </div>
            )}

            <button
              type="submit"
              disabled={profileForm.formState.isSubmitting || loadingProfile}
              className="mt-6 w-full bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white font-semibold py-3 rounded-lg transition-colors"
            >
              {profileForm.formState.isSubmitting ? 'Salvando...' : 'Salvar Alterações'}
            </button>
          </form>

          {/* Change password form */}
          <form onSubmit={passwordForm.handleSubmit(onChangePassword)} className="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h2 className="text-xs font-bold text-slate-500 uppercase tracking-widest mb-5">
              Alterar Senha
            </h2>
            <div className="space-y-3">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Senha Atual</label>
                <input type="password" placeholder="••••••••" {...passwordForm.register('current_password')} className={inputClass} />
                {passwordForm.formState.errors.current_password && (
                  <p className={errorClass}>{passwordForm.formState.errors.current_password.message}</p>
                )}
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Nova Senha</label>
                <input type="password" placeholder="••••••••" {...passwordForm.register('new_password')} className={inputClass} />
                {passwordForm.formState.errors.new_password && (
                  <p className={errorClass}>{passwordForm.formState.errors.new_password.message}</p>
                )}
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Confirmar Nova Senha</label>
                <input type="password" placeholder="••••••••" {...passwordForm.register('confirm_password')} className={inputClass} />
                {passwordForm.formState.errors.confirm_password && (
                  <p className={errorClass}>{passwordForm.formState.errors.confirm_password.message}</p>
                )}
              </div>
            </div>
            <button
              type="submit"
              disabled={passwordForm.formState.isSubmitting}
              className="mt-5 w-full bg-slate-700 hover:bg-slate-800 disabled:opacity-60 text-white font-semibold py-3 rounded-lg transition-colors"
            >
              {passwordForm.formState.isSubmitting ? 'Alterando...' : 'Alterar Senha'}
            </button>
          </form>
        </div>

        {/* Right column — Integrations (mock, Sprint 8 adiada) */}
        <div className="space-y-4">
          <h2 className="text-xs font-bold text-slate-500 uppercase tracking-widest">
            Integrações e Conexões
          </h2>

          <div className="bg-white rounded-xl shadow-sm border border-gray-100 p-5">
            <div className="flex items-start gap-4">
              <div className="w-10 h-10 rounded-lg bg-white border border-gray-200 flex items-center justify-center shrink-0 shadow-sm">
                <svg viewBox="0 0 24 24" className="w-6 h-6">
                  <path fill="#EA4335" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                  <path fill="#4285F4" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                  <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                  <path fill="#34A853" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
                </svg>
              </div>
              <div className="flex-1">
                <p className="font-semibold text-slate-800">Google Meet</p>
                {googleConnected
                  ? <p className="text-xs text-teal-600 font-medium">Conectado</p>
                  : <p className="text-xs text-slate-500">Conectar Conta</p>}
                {googleConnected && (
                  <p className="text-xs text-slate-500 mt-1">
                    E-mail: <span className="text-teal-600">dr.ana@gmail.com</span>
                  </p>
                )}
                <button className="mt-3 px-4 py-2 bg-teal-600 hover:bg-teal-700 text-white text-sm font-medium rounded-lg transition-colors">
                  {googleConnected ? 'Desconectar' : 'Conectar'}
                </button>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-xl shadow-sm border border-gray-100 p-5">
            <div className="flex items-start gap-4">
              <div className="w-10 h-10 rounded-lg bg-blue-600 flex items-center justify-center shrink-0 shadow-sm">
                <LinkIcon size={18} className="text-white" />
              </div>
              <div className="flex-1">
                <p className="font-semibold text-slate-800">Microsoft Outlook Calendar</p>
                <p className="text-xs text-slate-500">
                  {outlookConnected ? 'Conectado' : 'Conectar Conta'}
                </p>
                <p className="text-xs text-slate-400 mt-1">Sincronize sua agenda do Outlook</p>
                <button className="mt-3 px-4 py-2 bg-teal-600 hover:bg-teal-700 text-white text-sm font-medium rounded-lg transition-colors">
                  {outlookConnected ? 'Desconectar' : 'Conectar'}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
