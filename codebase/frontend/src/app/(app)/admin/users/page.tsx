'use client';

import { useState, useEffect } from 'react';
import { Plus, Trash2, X } from 'lucide-react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { api } from '@/lib/api';
import type { User } from '@/types';

const schema = z.object({
  name: z.string().min(2, 'Nome deve ter ao menos 2 caracteres'),
  email: z.string().email('E-mail inválido'),
  crp: z.string().regex(/^\d{2}\/\d{1,6}$/, 'CRP inválido (ex: 06/123456)').or(z.literal('')),
  specialty: z.string().optional(),
  password: z.string().min(6, 'A senha deve ter ao menos 6 caracteres'),
});

type FormData = z.infer<typeof schema>;

export default function AdminUsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);

  const {
    register,
    handleSubmit,
    reset,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({ resolver: zodResolver(schema) });

  useEffect(() => {
    api.get<User[]>('/api/admin/users')
      .then(setUsers)
      .finally(() => setLoading(false));
  }, []);

  function openModal() {
    reset();
    setShowModal(true);
  }

  async function onSubmit(data: FormData) {
    try {
      const newUser = await api.post<User>('/api/admin/users', data);
      setUsers((prev) => [...prev, newUser]);
      setShowModal(false);
    } catch {
      setError('root', { message: 'Erro ao criar psicólogo. Verifique os dados e tente novamente.' });
    }
  }

  async function handleDelete(id: string) {
    if (!confirm('Remover este psicólogo?')) return;
    await api.delete(`/api/admin/users/${id}`);
    setUsers((prev) => prev.filter((u) => u.id !== id));
  }

  const inputClass = 'w-full px-3 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500';
  const errorClass = 'text-red-500 text-xs mt-1';

  return (
    <div className="p-6 max-w-5xl mx-auto">
      <div className="flex items-center justify-between mb-6">
        <div>
          <h1 className="text-2xl font-bold text-slate-800">Psicólogos Cadastrados</h1>
          <p className="text-sm text-slate-500 mt-1">Gerenciamento de usuários do sistema</p>
        </div>
        <button
          onClick={openModal}
          className="flex items-center gap-2 bg-teal-600 hover:bg-teal-700 text-white text-sm font-medium px-4 py-2.5 rounded-lg transition-colors"
        >
          <Plus size={16} />
          Novo Psicólogo
        </button>
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
        {loading ? (
          <div className="p-8 space-y-3">
            {[...Array(3)].map((_, i) => (
              <div key={i} className="h-10 bg-gray-100 rounded animate-pulse" />
            ))}
          </div>
        ) : (
          <table className="w-full text-sm">
            <thead className="bg-gray-50 border-b border-gray-100">
              <tr>
                {['Nome', 'CRP', 'Especialidade', 'E-mail', 'Cadastro', ''].map((h) => (
                  <th key={h} className="text-left px-5 py-3 text-xs font-semibold text-slate-500 uppercase tracking-wider">{h}</th>
                ))}
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-50">
              {users.map((user) => (
                <tr key={user.id} className="hover:bg-gray-50 transition-colors">
                  <td className="px-5 py-4 font-medium text-slate-800">{user.name}</td>
                  <td className="px-5 py-4 text-slate-600">{user.crp || '—'}</td>
                  <td className="px-5 py-4 text-slate-600">{user.specialty || '—'}</td>
                  <td className="px-5 py-4 text-slate-600">{user.email}</td>
                  <td className="px-5 py-4 text-slate-400">
                    {new Date(user.created_at).toLocaleDateString('pt-BR')}
                  </td>
                  <td className="px-5 py-4">
                    <button
                      onClick={() => handleDelete(user.id)}
                      className="p-1.5 text-slate-400 hover:text-red-500 hover:bg-red-50 rounded transition-colors"
                    >
                      <Trash2 size={15} />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>

      {showModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40 px-4">
          <div className="bg-white rounded-2xl shadow-xl w-full max-w-md p-6">
            <div className="flex items-center justify-between mb-5">
              <h2 className="text-lg font-bold text-slate-800">Novo Psicólogo</h2>
              <button
                onClick={() => setShowModal(false)}
                className="p-1.5 hover:bg-gray-100 rounded-lg transition-colors"
              >
                <X size={18} className="text-slate-500" />
              </button>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Nome Completo</label>
                <input type="text" placeholder="Dr. João da Silva" {...register('name')} className={inputClass} />
                {errors.name && <p className={errorClass}>{errors.name.message}</p>}
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">E-mail</label>
                <input type="email" placeholder="dr.joao@email.com" {...register('email')} className={inputClass} />
                {errors.email && <p className={errorClass}>{errors.email.message}</p>}
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">CRP</label>
                <input type="text" placeholder="06/000000" {...register('crp')} className={inputClass} />
                {errors.crp && <p className={errorClass}>{errors.crp.message}</p>}
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Especialidade</label>
                <input type="text" placeholder="Psicologia Clínica" {...register('specialty')} className={inputClass} />
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-1">Senha Provisória</label>
                <input type="password" placeholder="••••••••" {...register('password')} className={inputClass} />
                {errors.password && <p className={errorClass}>{errors.password.message}</p>}
              </div>

              {errors.root && <p className="text-red-500 text-xs">{errors.root.message}</p>}

              <div className="flex gap-3 pt-2">
                <button
                  type="button"
                  onClick={() => setShowModal(false)}
                  className="flex-1 px-4 py-2.5 border border-gray-200 rounded-lg text-sm text-slate-600 hover:bg-gray-50 transition-colors"
                >
                  Cancelar
                </button>
                <button
                  type="submit"
                  disabled={isSubmitting}
                  className="flex-1 bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white text-sm font-medium px-4 py-2.5 rounded-lg transition-colors"
                >
                  {isSubmitting ? 'Criando...' : 'Criar Psicólogo'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
