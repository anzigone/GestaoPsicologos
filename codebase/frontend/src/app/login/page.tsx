'use client';

import { useState } from 'react';
import Image from 'next/image';
import { useRouter } from 'next/navigation';
import { Mail, Lock } from 'lucide-react';
import { api } from '@/lib/api';
import type { LoginResponse } from '@/types';

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      const data = await api.post<LoginResponse>('/api/auth/login', { email, password });
      document.cookie = `auth_token=${data.token}; path=/; max-age=${24 * 60 * 60}; SameSite=Strict`;
      router.push('/');
    } catch {
      setError('E-mail ou senha inválidos. Tente novamente.');
    } finally {
      setLoading(false);
    }
  }

  return (
    <div
      className="min-h-screen flex items-center justify-center px-4"
      style={{
        background: 'linear-gradient(135deg, #b2dfdb 0%, #e0f2f1 40%, #cde8e5 100%)',
      }}
    >
      <div className="bg-white rounded-2xl shadow-lg w-full max-w-sm px-8 py-10">
        <div className="flex flex-col items-center mb-6">
          <div className="mb-4">
            <Image
              src="/simbolo-psicologia.jpg"
              alt="Símbolo da Psicologia"
              width={80}
              height={80}
              className="rounded-full object-cover"
              onError={() => {}}
            />
          </div>
          <h1 className="text-2xl font-bold text-slate-800">Gestão Psicológica</h1>
          <p className="text-slate-500 text-sm mt-1">Login de Acesso</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-slate-700 mb-1">Email</label>
            <div className="relative">
              <Mail size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
              <input
                type="email"
                placeholder="user@email.com"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full pl-9 pr-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                required
              />
            </div>
          </div>

          <div>
            <div className="flex items-center justify-between mb-1">
              <label className="block text-sm font-medium text-slate-700">Senha</label>
              <button type="button" className="text-xs text-teal-600 hover:underline">
                esqueci a senha?
              </button>
            </div>
            <div className="relative">
              <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
              <input
                type="password"
                placeholder="••••••••"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full pl-9 pr-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                required
              />
            </div>
          </div>

          {error && <p className="text-red-500 text-xs">{error}</p>}

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white font-semibold py-3 rounded-lg transition-colors mt-2"
          >
            {loading ? 'Entrando...' : 'Entrar'}
          </button>
        </form>

        <div className="mt-6 text-center space-y-2">
          <p className="text-sm text-slate-500">
            <button className="hover:underline">Esqueceu a senha?</button>
          </p>
          <p className="text-sm text-slate-500">
            <button className="hover:underline">Criar Nova Conta</button>
          </p>
        </div>
      </div>
    </div>
  );
}
