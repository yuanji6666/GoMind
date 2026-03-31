import { useState, useEffect, useCallback } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { UserPlus, Eye, EyeOff, BrainCircuit, Mail } from 'lucide-react'
import toast from 'react-hot-toast'
import { register, sendCaptcha } from '@/api/gin'
import { CODE_SUCCESS } from '@/api/types'
import { useAuthStore } from '@/store/useAuthStore'

export default function RegisterPage() {
  const navigate = useNavigate()
  const setAuth = useAuthStore((s) => s.setAuth)
  const [email, setEmail] = useState('')
  const [captcha, setCaptcha] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPwd, setConfirmPwd] = useState('')
  const [showPwd, setShowPwd] = useState(false)
  const [submitting, setSubmitting] = useState(false)
  const [countdown, setCountdown] = useState(0)

  useEffect(() => {
    if (countdown <= 0) return
    const timer = setTimeout(() => setCountdown(countdown - 1), 1000)
    return () => clearTimeout(timer)
  }, [countdown])

  const handleSendCaptcha = useCallback(async () => {
    if (!email.trim()) {
      toast.error('请先输入邮箱')
      return
    }
    try {
      const res = await sendCaptcha({ email: email.trim() })
      if (res.status_code === CODE_SUCCESS) {
        toast.success('验证码已发送，请查看邮箱')
        setCountdown(60)
      } else {
        toast.error(res.status_msg || '发送失败')
      }
    } catch {
      toast.error('网络错误')
    }
  }, [email])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!email.trim() || !captcha.trim() || !password) {
      toast.error('请填写所有必填项')
      return
    }
    if (password !== confirmPwd) {
      toast.error('两次密码不一致')
      return
    }
    if (password.length < 6) {
      toast.error('密码长度至少6位')
      return
    }
    setSubmitting(true)
    try {
      const res = await register({ email: email.trim(), captcha: captcha.trim(), password })
      if (res.status_code === CODE_SUCCESS && res.token) {
        setAuth(res.token)
        toast.success('注册成功')
        navigate('/', { replace: true })
      } else {
        toast.error(res.status_msg || '注册失败')
      }
    } catch {
      toast.error('网络错误，请稍后重试')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 px-4">
      <div className="w-full max-w-md">
        <div className="text-center mb-8">
          <div className="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-indigo-600 mb-4">
            <BrainCircuit className="w-9 h-9 text-white" />
          </div>
          <h1 className="text-3xl font-bold text-white tracking-tight">GopherAI</h1>
          <p className="text-slate-400 mt-2">创建您的账号</p>
        </div>

        <form
          onSubmit={handleSubmit}
          className="bg-white/5 backdrop-blur-xl border border-white/10 rounded-2xl p-8 space-y-5"
        >
          <h2 className="text-xl font-semibold text-white text-center">注册</h2>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-1.5">邮箱</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full px-4 py-2.5 bg-white/5 border border-white/10 rounded-lg text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition"
              placeholder="your@email.com"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-1.5">验证码</label>
            <div className="flex gap-3">
              <input
                type="text"
                value={captcha}
                onChange={(e) => setCaptcha(e.target.value)}
                className="flex-1 px-4 py-2.5 bg-white/5 border border-white/10 rounded-lg text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition"
                placeholder="6位验证码"
                maxLength={6}
              />
              <button
                type="button"
                disabled={countdown > 0}
                onClick={handleSendCaptcha}
                className="flex items-center gap-1.5 px-4 py-2.5 bg-indigo-600/80 hover:bg-indigo-500 disabled:opacity-40 disabled:cursor-not-allowed text-white text-sm font-medium rounded-lg transition whitespace-nowrap"
              >
                <Mail className="w-4 h-4" />
                {countdown > 0 ? `${countdown}s` : '发送'}
              </button>
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-1.5">密码</label>
            <div className="relative">
              <input
                type={showPwd ? 'text' : 'password'}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full px-4 py-2.5 bg-white/5 border border-white/10 rounded-lg text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition pr-11"
                placeholder="至少6位密码"
              />
              <button
                type="button"
                onClick={() => setShowPwd(!showPwd)}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-white transition"
              >
                {showPwd ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-1.5">确认密码</label>
            <input
              type="password"
              value={confirmPwd}
              onChange={(e) => setConfirmPwd(e.target.value)}
              className="w-full px-4 py-2.5 bg-white/5 border border-white/10 rounded-lg text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition"
              placeholder="再次输入密码"
            />
          </div>

          <button
            type="submit"
            disabled={submitting}
            className="w-full flex items-center justify-center gap-2 py-2.5 bg-indigo-600 hover:bg-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed text-white font-medium rounded-lg transition"
          >
            <UserPlus className="w-4 h-4" />
            {submitting ? '注册中...' : '注册'}
          </button>

          <p className="text-center text-sm text-slate-400">
            已有账号？{' '}
            <Link to="/login" className="text-indigo-400 hover:text-indigo-300 font-medium transition">
              立即登录
            </Link>
          </p>
        </form>
      </div>
    </div>
  )
}
